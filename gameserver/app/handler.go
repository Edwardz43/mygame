package gameserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// var addr = flag.String("addr", ":8090", "http service address")

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func errHandle(err error) {
	if err != nil {
		return
	}
}

var isGaming bool

var conn *websocket.Conn

var hub *Hub

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
	r.LoadHTMLFiles("../resource/home.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.Static("/static", "../resource")

	r.GET("/ws", serveWebsocket)

	log.Println("listen http://localhost:8090")
	r.Run(":8090")
}

func startGame(hub *Hub) {
	result := make(chan *GameResult)
	go StartGame(result)
	for {
		r, err := json.Marshal(<-result)
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
	hub = newHub()
	go hub.run()
	go startGame(hub)
	serve()
}
