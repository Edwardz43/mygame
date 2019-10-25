package gameserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/gameserver/app/gamelogic"
	"github.com/Edwardz43/mygame/gameserver/app/service"
	"github.com/Edwardz43/mygame/gameserver/lib/log"
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

// type loginInfo struct {
// 	Run       int64 `json:"run"`
// 	Inn       int   `json:"inn"`
// 	Status    int8  `json:"status"`
// 	Countdown int8  `json:"countdown"`
// }

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
	hub               *Hub
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
	command           chan *Data
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
	hub = newHub()
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

	client := &Client{
		memberID: uint(id),
		hub:      hub,
		conn:     conn,
		send:     make(chan []byte, 256),
	}

	hub.register <- client

	nowRun, nowInn, nowStatus, nowCountdown, _ := lobbyService.GetLatest(1)

	latestResult, err := gameResultService.GetLatest(int8(gameBase.GetGameID()), run, inn-1)
	errHandle(err)

	format := "{\"Run\":%d, \"Inn\":%d, \"Status\":%d, \"Countdown\":%d, \"Result\":%s}"

	d, err := json.Marshal(Data{
		Event:   "200",
		Message: fmt.Sprintf(format, nowRun, nowInn, nowStatus, nowCountdown, latestResult),
	})

	errHandle(err)

	hub.send <- &PersonalMessage{
		client:  client,
		message: d,
	}

	command = make(chan *Data)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	for {
		select {
		case c := <-command:
			logger.Printf("COMMAND : [%v], DATA: [%v]", c.Event, c.Message)

			switch c.Event {
			case "300": // get table status
				//TODO
				break
			case "301": // bet

				msg, err := bet(client.memberID, c.Message)

				if err != nil {
					//TODO
					d, err = json.Marshal(Data{
						Event:   "301",
						Message: err.Error(),
					})
				} else {
					d, err = json.Marshal(Data{
						Event:   "301",
						Message: msg,
					})
				}

				hub.send <- &PersonalMessage{
					client:  client,
					message: d,
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

	// Logger.Println("listen http://localhost:8090")
	engine.Run(":8090")
}

func start(hub *Hub, gb gamelogic.GameBase) {

	gameResult = new(gamelogic.GameResult)

	gameBase = gb
	// for {
	go gameBase.StartGame()

	run, inn, status, oldCountdown, _ = lobbyService.GetLatest(int(gamelogic.Dice))

	// if err != nil {
	// 	panic(err)
	// }

	if run == 0 {
		i, _ := strconv.Atoi(time.Now().Format("20060102"))
		run = int64(i)
	}

	if inn == 0 {
		inn = 1
	}

	switch GameStatus(status) {
	case NewInn:
		//TODO
		newInn()
		break
	case Showdown:
		//TODO
		showDown()
		break
	case Settlement:
		//TODO
		settlement()
		break
	case Intermission:
		//TODO
		break
	case Maintain:
		//TODO
		break
	default:
		//TODO
	}
}

// newRun 新輪
func newRun() {
	// Logger.Printf("[%s] : [%s]", "hanlder", "newRun")
	runOld, _, _, _, err := lobbyService.GetLatest(int(gameBase.GetGameID()))
	errHandle(err)
	runNow, _ := strconv.Atoi(time.Now().Format("20060102"))
	if runOld != int64(runNow) {
		run = int64(runNow)
		inn = 0
		lobbyService.Update(int(gameBase.GetGameID()), run, 1, int(NewInn))
	}
	newInn()
}

// newInn 新局
func newInn() {
	inn++

	detail := gameBase.NewGame()

	lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(NewInn))

	gameResult.Run = run
	gameResult.Inn = inn
	gameResult.GameType = gameBase.GetGameID()
	gameResult.GameDetail = detail

	newRun := Data{
		Event:   "201",
		Message: duration.String()[0:2],
	}

	d, err := json.Marshal(newRun)

	errHandle(err)

	hub.broadcast <- d

	// time.AfterFunc(duration, showDown)

	ticker := time.NewTicker(time.Second)

	var count int8

	if oldCountdown == 0 {
		count = 20
	} else {
		count = oldCountdown
		oldCountdown = 0
	}

	for count > -1 {
		select {
		case <-ticker.C:
			err := lobbyService.Countdown(int(gameBase.GetGameID()), int8(count))
			errHandle(err)
			count--
		}
	}
	showDown()
}

// showDown 開牌
func showDown() {

	lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(Showdown))

	detail, _ := json.Marshal(gameResult.GameDetail)

	go func() {
		m, err := gameResultService.
			AddNewOne(int8(gameResult.GameType), gameResult.Run, gameResult.Inn, string(detail), 0)
		errHandle(err)
		logger.Printf("[%s] : [%s] message [%s]", "GameResultService", "AddNewOne", m)
	}()

	r, err := json.Marshal(gameResult)
	errHandle(err)
	data := Data{
		Event:   "202",
		Message: string(r),
	}
	d, err := json.Marshal(data)
	errHandle(err)
	hub.broadcast <- d

	time.AfterFunc(showDownTime, settlement)
}

// settlement 結算
func settlement() {

	lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(Settlement))

	data := Data{
		Event:   "203",
		Message: "Settling",
	}
	d, err := json.Marshal(data)
	errHandle(err)
	hub.broadcast <- d

	time.AfterFunc(showDownTime, newRun)
}

func intermission() {}

func maintain() {}

func bet(memberID uint, msg string) (string, error) {

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

// Startup starts process.
func Startup() {
	// isGaming = false
	go hub.run()
	go start(hub, &gamelogic.DiceGame{})
	serve()
}
