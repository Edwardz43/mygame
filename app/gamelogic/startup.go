package gamelogic

import (
	"strconv"
	"time"
)

// GameStatus ...
type GameStatus int8

// COMMAND ...
type COMMAND int

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
	gameResult   *GameResult
	gameBase     GameBase
	run          int64
	inn          int
	status       int8
	oldCountdown int8
)

func errHandle(err error) {
	if err == nil {
		return
	}
	logger.Printf("ERROR : [%v]", err)
}

// Start ...
func Start(gb GameBase) {

	gameResult = new(GameResult)

	gameBase = gb
	// for {
	go gameBase.StartGame()

	// run, inn, status, oldCountdown, _ = lobbyService.GetLatest(int(gameBase.GetGameID()))

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
	// runOld, _, _, _, err := lobbyService.GetLatest(int(gameBase.GetGameID()))
	// errHandle(err)
	// runNow, _ := strconv.Atoi(time.Now().Format("20060102"))
	// if runOld != int64(runNow) {
	// 	run = int64(runNow)
	// 	inn = 0
	// 	lobbyService.Update(int(gameBase.GetGameID()), run, 1, int(NewInn))
	// }
	newInn()
}

// newInn 新局
func newInn() {
	inn++

	detail := gameBase.NewGame()

	// lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(NewInn))

	gameResult.Run = run
	gameResult.Inn = inn
	gameResult.GameType = gameBase.GetGameID()
	gameResult.GameDetail = detail

	// newRun := Data{
	// 	Event:   "201",
	// 	Message: duration.String()[0:2],
	// }

	// d, err := json.Marshal(newRun)

	// errHandle(err)

	// hub.broadcast <- d

	// time.AfterFunc(duration, showDown)

	// ticker := time.NewTicker(time.Second)

	// var count int8

	// if oldCountdown == 0 {
	// 	count = 20
	// } else {
	// 	count = oldCountdown
	// 	oldCountdown = 0
	// }

	// for count > -1 {
	// 	select {
	// 	case <-ticker.C:
	// 		err := lobbyService.Countdown(int(gameBase.GetGameID()), int8(count))
	// 		errHandle(err)
	// 		count--
	// 	}
	// }
	showDown()
}

// showDown 開牌
func showDown() {
	/*
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

		time.AfterFunc(showDownTime, settlement)*/
}

// settlement 結算
func settlement() {
	/*
		lobbyService.Update(int(gameBase.GetGameID()), run, inn, int(Settlement))

		data := Data{
			Event:   "203",
			Message: "Settling",
		}
		d, err := json.Marshal(data)
		errHandle(err)
		hub.broadcast <- d

		time.AfterFunc(showDownTime, newRun)
	*/
}

func intermission() {}

func maintain() {}
