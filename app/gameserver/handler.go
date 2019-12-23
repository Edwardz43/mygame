package gameserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/app/gamelogic"
	"github.com/Edwardz43/mygame/app/lib/log"
	"github.com/Edwardz43/mygame/app/service"
	socket "github.com/Edwardz43/mygame/app/websocket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GameStatus ...
type GameStatus int8

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
	NewInn GameStatus = iota + 1
	Showdown
	Settlement
	Intermission
	Maintain
)

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
	hub               *socket.Hub
	bettingService    *service.BettingService
	gameResultService *service.GameResultService
	lobbyService      *service.LobbyService
	memberService     *service.MemberService
	gameResult        *gamelogic.GameResult
	gameBase          gamelogic.GameBase
	run               int64
	inn               int
	status            int8
	oldCountdown      int8
	command           chan *socket.Data
	logger            *log.Logger
	tokenMap          map[string]bool
	duration          = time.Second * 20
	showDownTime      = time.Second * 3
	settlementTime    = time.Second * 5
)

func init() {
	logger = log.Create("gameserver")
	gameResultService = service.GetGameResultInstance()
	lobbyService = service.GetLobbyInstance()
	memberService = service.GetLoginInstance()
	bettingService = service.GetBettingInstance()
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

	nowRun, nowInn, nowStatus, nowCountdown, _ := lobbyService.GetLatest(1)

	latestResult, err := gameResultService.GetLatest(int8(gameBase.GetGameID()), run, inn-1)
	errHandle(err)

	format := "{\"GameID\":%d,\"Run\":%d, \"Inn\":%d, \"Status\":%d, \"Countdown\":%d, \"Result\":%v}"

	d, err := json.Marshal(socket.Data{
		Event:   "200",
		Message: fmt.Sprintf(format, int8(gameBase.GetGameID()), nowRun, nowInn, nowStatus, nowCountdown, latestResult),
	})

	errHandle(err)

	hub.Send <- &socket.PersonalMessage{
		Client:  client,
		Message: d,
	}

	command = make(chan *socket.Data)

	client.Listen(command)

	for {
		select {
		case c := <-command:
			logger.Printf("COMMAND : [%v], DATA: [%v]", c.Event, c.Message)

			switch c.Event {
			case "300": // get table status
				//TODO
				break
			case "301": // bet

				msg, err := bet(client.MemberID, c.Message)

				if err != nil {
					//TODO
					d, err = json.Marshal(socket.Data{
						Event:   "301",
						Message: err.Error(),
					})
				} else {
					d, err = json.Marshal(socket.Data{
						Event:   "301",
						Message: msg,
					})
				}

				hub.Send <- &socket.PersonalMessage{
					Client:  client,
					Message: d,
				}

				break
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
func Startup(gb gamelogic.GameBase) {
	// isGaming = false
	go hub.Run()
	go start(gb)
	serve()
}
