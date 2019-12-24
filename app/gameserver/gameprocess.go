package gameserver

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/app/lib/log"
	"github.com/Edwardz43/mygame/app/websocket"

	"github.com/Edwardz43/mygame/app/gamelogic"

	"github.com/Edwardz43/mygame/app/service"
)

// GameProcess creates a game process instance
type GameProcess struct {
	Hub      net.Hub
	GameBase gamelogic.GameBase
}

// GameStatus ...
type GameStatus int8

const (
	NewInn GameStatus = iota + 1
	Showdown
	Settlement
	Intermission
	Maintain
)

var (
	gameResultService *service.GameResultService
	gameResult        *gamelogic.GameResult
	run               int64
	inn               int
	status            int8
	oldCountdown      int8
	command           chan *websocket.Data
	logger            *log.Logger
	lobbyService      *service.LobbyService
	gameBase          gamelogic.GameBase
	hub               *websocket.Hub
	duration          = time.Second * 20
	showDownTime      = time.Second * 3
	settlementTime    = time.Second * 5
)

func errHandle(err error) {
	if err == nil {
		return
	}
	logger.Printf("ERROR : [%v]", err)
}

func init() {
	logger = log.Create("gameserver")
	gameResultService = service.GetGameResultInstance()
	lobbyService = service.GetLobbyInstance()
	gameResult = new(gamelogic.GameResult)
}

// Start starts game process
func (p *GameProcess) Start() {
	gameBase = p.GameBase
	go gameBase.StartGame()

	run, inn, status, oldCountdown, _ = lobbyService.GetLatest(int(gamelogic.Dice))
	logger.Printf(fmt.Sprintf("INN : %d", inn))

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
		inn = 1
		lobbyService.Update(int(gameBase.GetGameID()), run, 1, int(NewInn))
	}
	newInn()
}

// newInn 新局
func newInn() {

	detail := gameBase.NewGame()

	lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(NewInn))
	logger.Printf(fmt.Sprintf("newInn INN : %d", inn))
	gameResult.Run = run
	gameResult.Inn = inn

	gameResult.GameType = gameBase.GetGameID()
	gameResult.GameDetail = detail

	newRun := websocket.Data{
		Event:   "201",
		Message: fmt.Sprintf("{\"game_id\":%d,\"run\":%d, \"inn\":%d, \"countdown\":%s}", gameBase.GetGameID(), run, inn, duration.String()[0:2]),
	}

	d, err := json.Marshal(newRun)

	errHandle(err)

	hub.Broadcast <- d

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
			logger.Printf("countdown : %d", count)

			newRun := websocket.Data{
				Event:   "205",
				Message: fmt.Sprintf("{\"game_id\":%d,\"run\":%d, \"inn\":%d, \"countdown\":%d}", gameBase.GetGameID(), run, inn, count),
			}

			d, err := json.Marshal(newRun)

			errHandle(err)

			hub.Broadcast <- d

			err = lobbyService.Countdown(int(gameBase.GetGameID()), int8(count))
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
	logger.Printf(fmt.Sprintf("INN : %d", inn))

	r, err := json.Marshal(gameResult)
	errHandle(err)
	data := websocket.Data{
		Event:   "202",
		Message: string(r),
	}
	d, err := json.Marshal(data)
	errHandle(err)
	hub.Broadcast <- d

	time.AfterFunc(showDownTime, settlement)
}

// settlement 結算
func settlement() {

	lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(Settlement))

	data := websocket.Data{
		Event:   "203",
		Message: "Settling",
	}
	d, err := json.Marshal(data)
	errHandle(err)
	hub.Broadcast <- d
	logger.Printf(fmt.Sprintf("INN : %d", inn))

	time.AfterFunc(showDownTime, newRun)

	inn++
}

func intermission() {}

func maintain() {}
