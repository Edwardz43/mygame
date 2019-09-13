package gameserver

import (
	"encoding/json"
	"log"
	"net/http"
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
	upGrader          = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
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
	result := make(chan *GameResult)
	go gb.StartGame(result)
	for {

		// lobbyService := service.GetLobbyInstance()

		// run, inn, status, err := lobbyService.GetLatest(int(Dice))

		// gameR := gb.NewGame()
		newRun := Data{
			Event:   "201",
			Message: duration.String()[0:2],
		}

		d, err := json.Marshal(newRun)

		errHandle(err)

		hub.broadcast <- d

		gameR := <-result

		detail, _ := json.Marshal(gameR.GameDetail)

		time.AfterFunc(duration, func() {

			go func() {
				m, err := gameResultService.
					AddNewOne(int8(gameR.GameType), gameR.Run, gameR.Inn, string(detail), 0)
				errHandle(err)
				log.Println(m)
			}()

			r, err := json.Marshal(gameR)
			errHandle(err)
			data := Data{
				Event:   "202",
				Message: string(r),
			}
			d, err := json.Marshal(data)
			errHandle(err)
			hub.broadcast <- d
		})
	}
}

// Startup starts process.
func Startup() {
	// isGaming = false
	gameResultService = service.GetGameResultInstance()
	hub = newHub()
	go hub.run()
	go start(hub, &DiceGame{})
	serve()
}
