package gameserver

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/app/lib/log"
	"github.com/Edwardz43/mygame/app/lib/nettool"

	"github.com/Edwardz43/mygame/app/gamelogic"

	"github.com/Edwardz43/mygame/app/service"
)

// GameProcess creates a game process instance
type GameProcess struct {
	Hub          nettool.Hub
	GameBase     gamelogic.GameBase
	gameResult   *gamelogic.GameResult
	run          int64
	inn          int
	status       int8
	oldCountdown int8
	duration     time.Duration
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
	command           chan *nettool.Data
	logger            *log.Logger
	lobbyService      *service.LobbyService
	gameBase          gamelogic.GameBase
	showDownTime      = time.Second * 3
	settlementTime    = time.Second * 3
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

}

// Start starts game process
func (p *GameProcess) Start() {
	gameBase = p.GameBase
	// logger.Printf("GameBase[%v]", gameBase.GetGameID())

	go gameBase.StartGame()
	p.gameResult = new(gamelogic.GameResult)

	gameID := p.GameBase.GetGameID()
	// logger.Printf("GameID[%d] start", gameID)

	switch gameID {
	case gamelogic.Dice:
		p.duration = time.Second * 30
	case gamelogic.DragonTiger:
		p.duration = time.Second * 20
	}

	p.run, p.inn, p.status, p.oldCountdown, _ = lobbyService.GetLatest(int(gameID))
	logger.Printf(fmt.Sprintf("gameID[%d], run[%d], inn[%d], status[%d], cd[%d]", gameID, p.run, p.inn, p.status, p.oldCountdown))

	// if err != nil {
	// 	panic(err)
	// }

	if p.run == 0 {
		i, _ := strconv.Atoi(time.Now().Format("20060102"))
		p.run = int64(i)
	}

	if p.inn == 0 {
		p.inn = 1
	}

	// logger.Printf("[%s] : [%s]", "Start", status)

	switch GameStatus(p.status) {
	case NewInn:
		//TODO
		p.newInn()
		break
	case Showdown:
		//TODO
		p.showDown()
		break
	case Settlement:
		//TODO
		p.settlement()
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
func (p *GameProcess) newRun() {
	logger.Printf("[%s] : [%s]", "hanlder", "newRun")
	runOld, _, _, _, err := lobbyService.GetLatest(int(gameBase.GetGameID()))
	errHandle(err)
	runNow, _ := strconv.Atoi(time.Now().Format("20060102"))
	if runOld != int64(runNow) {
		p.run = int64(runNow)
		p.inn = 1
		lobbyService.Update(int(gameBase.GetGameID()), p.run, 1, int(NewInn))
	}
	p.newInn()
}

// newInn 新局
func (p *GameProcess) newInn() {

	detail := p.GameBase.NewGame()

	lobbyService.Update(int(p.GameBase.GetGameID()), p.run, p.inn, int(NewInn))
	logger.Printf(fmt.Sprintf("GameID[%d] NewInn: %d", int(gameBase.GetGameID()), p.inn))
	p.gameResult.Run = p.run
	p.gameResult.Inn = p.inn

	p.gameResult.GameType = p.GameBase.GetGameID()
	p.gameResult.GameDetail = detail

	newRun := nettool.Data{
		Event:   "201",
		Message: fmt.Sprintf("{\"game_type\":%d,\"run\":%d, \"inn\":%d, \"countdown\":%d}", p.GameBase.GetGameID(), p.run, p.inn, int8(p.duration.Seconds())),
	}

	d, err := json.Marshal(newRun)

	errHandle(err)

	p.Hub.Broadcast <- d

	ticker := time.NewTicker(time.Second)

	var count int8

	if p.oldCountdown == 0 {
		count = int8(p.duration.Seconds())
	} else {
		count = p.oldCountdown
		p.oldCountdown = 0
	}

	for count > -1 {
		select {
		case <-ticker.C:
			logger.Printf("GameID[%d] countdown : %d", int(p.GameBase.GetGameID()), count)

			newRun := nettool.Data{
				Event:   "205",
				Message: fmt.Sprintf("{\"game_type\":%d,\"run\":%d, \"inn\":%d, \"countdown\":%d}", p.GameBase.GetGameID(), p.run, p.inn, count),
			}

			d, err := json.Marshal(newRun)

			errHandle(err)

			p.Hub.Broadcast <- d

			err = lobbyService.Countdown(int(p.GameBase.GetGameID()), int8(count))
			errHandle(err)
			count--
		}
	}
	p.showDown()
}

// showDown 開牌
func (p *GameProcess) showDown() {

	lobbyService.Update(int(p.GameBase.GetGameID()), p.run, p.inn, int(Showdown))

	detail, _ := json.Marshal(p.gameResult.GameDetail)

	go func() {
		m, err := gameResultService.
			AddNewOne(int8(p.gameResult.GameType), p.gameResult.Run, p.gameResult.Inn, string(detail), 0)
		errHandle(err)
		logger.Printf("[%s] : [%s] message [%s]", "GameResultService", "AddNewOne", m)
	}()
	logger.Printf(fmt.Sprintf("INN : %d", p.inn))

	r, err := json.Marshal(p.gameResult)
	errHandle(err)
	data := nettool.Data{
		Event:   "202",
		Message: string(r),
	}
	d, err := json.Marshal(data)
	errHandle(err)
	p.Hub.Broadcast <- d

	time.AfterFunc(showDownTime, p.settlement)
}

// settlement 結算
func (p *GameProcess) settlement() {

	lobbyService.Update(int(p.GameBase.GetGameID()), p.run, p.inn, int(Settlement))

	data := nettool.Data{
		Event:   "203",
		Message: fmt.Sprintf("{\"game_type\":%d}", p.GameBase.GetGameID()),
	}
	d, err := json.Marshal(data)
	errHandle(err)
	p.Hub.Broadcast <- d
	logger.Printf(fmt.Sprintf("INN : %d", p.inn))

	time.AfterFunc(showDownTime, p.newRun)

	p.inn++
}

func (p *GameProcess) intermission() {}

func (p *GameProcess) maintain() {}
