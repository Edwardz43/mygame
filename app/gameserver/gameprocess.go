package gameserver

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/app/gamelogic"
	socket "github.com/Edwardz43/mygame/app/websocket"
)

//TODO

func start(gb gamelogic.GameBase) {

	gameResult = new(gamelogic.GameResult)

	gameBase = gb
	// for {
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

	newRun := socket.Data{
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
	logger.Printf(fmt.Sprintf("INN : %d", inn))

	r, err := json.Marshal(gameResult)
	errHandle(err)
	data := socket.Data{
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

	data := socket.Data{
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
