package gameserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/gameserver/app/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func errHandle(err error) {
	if err != nil {
		return
	}
}

// var addr = flag.String("addr", ":8090", "http service address")
var (
	isGaming          bool
	conn              *websocket.Conn
	hub               *Hub
	gameResultService *service.GameResultService
	lobbyService      *service.LobbyService
	upGrader          = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	gameResult *GameResult
	gameBase   GameBase
	run        int64
	inn        int
	status     int8
)

func serveWebsocket(c *gin.Context) {
	// flag.Parse()
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	errHandle(err)
	client := &Client{
		ID:   1,
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func serve() {
	r := gin.Default()
	r.LoadHTMLFiles("./resource/home.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.Static("/static", "./resource")

	r.GET("/ws", serveWebsocket)

	log.Println("listen http://localhost:8090")
	r.Run(":8090")
}

func start(hub *Hub, gb GameBase) {

	gameResult = new(GameResult)

	gameBase = gb
	// for {
	go gameBase.StartGame()

	run, inn, status, _ = lobbyService.GetLatest(int(Dice))

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

// Startup starts process.
func Startup() {
	// isGaming = false
	gameResultService = service.GetGameResultInstance()
	lobbyService = service.GetLobbyInstance()
	hub = newHub()
	go hub.run()
	go start(hub, &DiceGame{})
	serve()
}

func newRun() {
	log.Printf("[%s] : [%s]", "hanlder", "newRun")
	runOld, _, _, err := lobbyService.GetLatest(int(gameBase.GetGameID()))
	errHandle(err)
	runNow, _ := strconv.Atoi(time.Now().Format("20060102"))
	if runOld != int64(runNow) {
		run = int64(runNow)
		inn = 0
		lobbyService.Update(int(gameBase.GetGameID()), run, 1, int(NewInn))
	}
	newInn()
}

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

	time.AfterFunc(duration, showDown)
}

func showDown() {

	lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(Showdown))

	detail, _ := json.Marshal(gameResult.GameDetail)

	go func() {
		m, err := gameResultService.
			AddNewOne(int8(gameResult.GameType), gameResult.Run, gameResult.Inn, string(detail), 0)
		errHandle(err)
		log.Printf("[%s] : [%s] message [%s]", "GameResultService", "AddNewOne", m)
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
