package gameserver

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

var inn = 1
var duration = time.Second * 10

// DiceGame ...
type DiceGame struct{}

// StartGame ...
func (d *DiceGame) StartGame(result chan *GameResult) {

	log.Println("Start Game")
	r := d.NewGame()
	result <- r
	for range time.Tick(duration) {
		r := d.NewGame()
		result <- r
	}
}

// NewGame ...
func (d *DiceGame) NewGame() *GameResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	run, _ := strconv.Atoi(time.Now().Format("20060102"))
	gr := GameResult{
		Run:      int64(run),
		Inn:      inn,
		GameType: Dice,
		GameDetail: DiceGameDetail{
			D1: r.Intn(6) + 1,
			D2: r.Intn(6) + 1,
			D3: r.Intn(6) + 1,
		},
	}
	inn++
	log.Println(gr)
	return &gr
}
