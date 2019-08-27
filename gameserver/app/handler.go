package gameserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/gameserver/app/service"
	"github.com/Edwardz43/mygame/gameserver/db"
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

func startGame(hub *Hub, gb GameBase) {
	result := make(chan *GameResult)
	go gb.StartGame(result)
	for {
		gameR := <-result

		detail, _ := json.Marshal(gameR.GameDetail)

		run, _ := strconv.Atoi(time.Now().Format("20060102") + fmt.Sprintf("%04d", gameR.Run))

		go func() {
			m, err := gameResultService.AddNewOne(int8(gameR.GameType), int64(run), string(detail), 0)
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
	}
}

// Start starts process.
func Start() {
	// isGaming = false
	gameResultService = &service.GameResultService{
		DbConn: db.Connect(),
	}
	hub = newHub()
	go hub.run()
	go startGame(hub, &DiceGame{})
	serve()
}
