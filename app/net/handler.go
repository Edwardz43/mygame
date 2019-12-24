package net

import (
	"encoding/json"
	socket "github.com/Edwardz43/mygame/app/lib/websocket"
	"net/http"
	"strconv"

	"github.com/Edwardz43/mygame/app/gameserver"

	"github.com/Edwardz43/mygame/app/gamelogic"
	"github.com/Edwardz43/mygame/app/lib/log"
	"github.com/Edwardz43/mygame/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// COMMAND ...
type COMMAND int

// betOrder from frontend
// {"game":1, "bet-area":"even", "amount":"100"}
type betOrder struct {
	Game    int8   `json:"game"`
	BetArea string `json:"bet-area"`
	Amount  int    `json:"amount"`
}

const (
	Register COMMAND = iota + 200
	NewRun
	ShowDown
	Result
	Bet
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	isGaming          bool
	engine            *gin.Engine
	conn              *websocket.Conn
	lobbyService      *service.LobbyService
	memberService     *service.MemberService
	gameResultService *service.GameResultService
	bettingService    *service.BettingService
	logger            *log.Logger
	tokenMap          map[string]bool
	hub               *socket.Hub
	command           chan *socket.Data
)

func init() {
	logger = log.Create("net")
	bettingService = service.GetBettingInstance()
	lobbyService = service.GetLobbyInstance()
	memberService = service.GetLoginInstance()
	hub = socket.NewHub()
	engine = gin.Default()
	tokenMap = make(map[string]bool)
}

func errHandle(err error) {
	if err == nil {
		return
	}
	logger.Printf("ERROR : [%v]", err)
}

// var addr = flag.String("addr", ":8090", "http service address")

func bet(memberID uint, run int64, inn int, msg string) (string, error) {

	logger.Printf("BETTING ID : [%v], data : [%v]", memberID, msg)
	// TODO

	var b betOrder

	err := json.Unmarshal([]byte(msg), &b)
	if err != nil {
		logger.Println("BETTING fail : json unmarshal")
		return "", err
	}

	var distinctID int

	switch b.BetArea {
	case "big":
		distinctID = 1
		break
	case "small":
		distinctID = 2
		break
	case "odd":
		distinctID = 3
		break
	case "even":
		distinctID = 4
		break
	}

	i, err := bettingService.AddNewOne(int8(b.Game), run, inn, int(memberID), distinctID, b.Amount)

	if err != nil {
		logger.Println("BETTING fail : BettingService")
		return "", err
	}

	logger.Println("BETTING ok")
	return i, nil

}

func serveWebsocket(c *gin.Context) {
	// flag.Parse()
	memberID := c.Query("memberID")
	id, _ := strconv.Atoi(memberID)
	// if tokenMap[token] {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	defer conn.Close()

	errHandle(err)

	client := &socket.Client{}
	client.Set(uint(id), hub, conn, make(chan []byte, 256))

	hub.Register <- client

	// nowRun, nowInn, nowStatus, nowCountdown, _ := lobbyService.GetLatest(1)

	// latestResult, err := gameResultService.GetLatest(int8(gameBase.GetGameID()), run, inn-1)
	// errHandle(err)

	// format := "{\"game_id\":%d,\"run\":%d, \"inn\":%d, \"status\":%d, \"countdown\":%d, \"result\":%v}"

	// d, err := json.Marshal(Data{
	// 	Event:   "200",
	// 	Message: fmt.Sprintf(format, int8(gameBase.GetGameID()), nowRun, nowInn, nowStatus, nowCountdown, latestResult),
	// })

	// errHandle(err)

	// hub.Send <- &PersonalMessage{
	// 	Client:  client,
	// 	Message: d,
	// }

	command = make(chan *socket.Data)

	client.Listen(command)

	for {
		select {
		case c := <-command:
			logger.Printf("COMMAND : [%v], DATA: [%v]", c.Event, c.Message)

			// var d []byte
			switch c.Event {
			case "300": // get table status
				//TODO
				break
				// case "301": // bet

				// 	msg, err := bet(client.MemberID, c.Message)

				// 	if err != nil {
				// 		//TODO
				// 		d, _ = json.Marshal(Data{
				// 			Event:   "301",
				// 			Message: err.Error(),
				// 		})
				// 	} else {
				// 		d, _ = json.Marshal(Data{
				// 			Event:   "301",
				// 			Message: msg,
				// 		})
				// 	}

				// 	hub.Send <- &PersonalMessage{
				// 		Client:  client,
				// 		Message: d,
				// 	}

				// 	break
			}

		}
	}
}

func serve() {
	// resource
	engine.Static("/static", "./resource")

	// index
	engine.GET("/", func(c *gin.Context) {
		engine.LoadHTMLFiles("./resource/index.html")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// login
	engine.POST("/login", login)

	// register
	engine.POST("/register", register)

	// game
	engine.GET("/game", func(c *gin.Context) {
		engine.LoadHTMLFiles("./resource/game.html")
		c.HTML(http.StatusOK, "game.html", nil)
	})

	engine.GET("/ws", serveWebsocket)

	engine.Run(":8090")
}

// Startup starts process.
func Startup() {
	// isGaming = false
	go hub.Run()

	dice := &gameserver.GameProcess{Hub: hub, GameBase: &gamelogic.DiceGame{}}

	go dice.GameBase.StartGame()

	// dt := &gameserver.GameProcess{Hub: hub, GameBase: &gamelogic.DragonTigerGame{}}

	// go dt.GameBase.StartGame()

	// go start(gb)
	serve()
}
