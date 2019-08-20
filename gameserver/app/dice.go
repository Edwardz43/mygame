package gameserver

import (
	"log"
	"math/rand"
	"time"
)

var run = 1

// DiceGame ...
type DiceGame struct{}

// StartGame ...
func (d *DiceGame) StartGame(result chan *GameResult) {
	log.Println("Start Game")
	// t := time.NewTicker(time.Second * 30)
	// for {
	// 	select {
	// 	case <-t.C:
	// 		log.Println("New Game")
	// 		r := NewGame()
	// 		// log.Println(r)
	// 		result <- r
	// 	default:
	// 		//
	// 	}
	// }
	for range time.Tick(time.Second * 10) {
		log.Println("New Game")
		r := d.NewGame()
		result <- r
	}
}

// NewGame ...
func (d *DiceGame) NewGame() *GameResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	gr := GameResult{
		Run:       run,
		Timestamp: time.Now(),
		GameType:  Dice,
		GameDetail: DiceGameDetail{
			D1: r.Intn(6) + 1,
			D2: r.Intn(6) + 1,
			D3: r.Intn(6) + 1,
		},
	}
	run++
	log.Println(gr)
	return &gr
}
