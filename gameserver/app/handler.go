package gameserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/gameserver/app/service"
	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/Edwardz43/mygame/gameserver/lib/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	isGaming bool

	engine   *gin.Engine
	conn     *websocket.Conn
	hub      *Hub
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	gameResultService *service.GameResultService
	lobbyService      *service.LobbyService
	memberService     *service.MemberService

	gameResult *GameResult
	gameBase   GameBase

	run     int64
	inn     int
	status  int8
	command chan *Data
	logger  *log.Logger
)

func init() {
	logger = log.Create("gameserver")
	gameResultService = service.GetGameResultInstance()
	lobbyService = service.GetLobbyInstance()
	memberService = service.GetLoginInstance()
	hub = newHub()
	engine = gin.Default()
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
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	errHandle(err)
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.register <- client

	msg := Data{
		Event:   "200",
		Message: "",
	}

	d, err := json.Marshal(msg)

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
		command := <-command

		switch command.Event {
		case "300":
			logger.Printf("COMMAND : [%v], DATA: [%v]", command.Event, command.Message)
		case "200": // register member to ws client
			logger.Printf("COMMAND : [%v], DATA: [%v]", command.Event, command.Message)
			m := new(models.Member)
			json.Unmarshal([]byte(command.Message), &m)
			client.memberID = m.ID
			logger.Printf("MEMBER : [%v]", client.memberID)
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

// newRun 新輪
func newRun() {
	// Logger.Printf("[%s] : [%s]", "hanlder", "newRun")
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

	time.AfterFunc(duration, showDown)
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

// Startup starts process.
func Startup() {
	// isGaming = false
	go hub.run()
	go start(hub, &DiceGame{})
	serve()
}
