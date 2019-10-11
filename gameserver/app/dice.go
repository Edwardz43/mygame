package gameserver

import (
	"math/rand"
	"time"
)

var duration = time.Second * 20
var showDownTime = time.Second * 3
var settlementTime = time.Second * 5

// DiceGameDetail ...
type DiceGameDetail struct {
	D1 int `json:"d1"`
	D2 int `json:"d2"`
	D3 int `json:"d3"`
}

// DiceGame ...
type DiceGame struct{}

// StartGame ...
func (d *DiceGame) StartGame() {
	Logger.Println("Start Game")
}

// NewGame ...
func (d *DiceGame) NewGame() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// run, _ := strconv.Atoi(time.Now().Format("20060102"))
	// gr := GameResult{
	// 	Run:      int64(run),
	// 	Inn:      inn,
	// 	GameType: Dice,
	// 	GameDetail: DiceGameDetail{
	// 		D1: r.Intn(6) + 1,
	// 		D2: r.Intn(6) + 1,
	// 		D3: r.Intn(6) + 1,
	// 	},
	// }
	// inn++
	detail := DiceGameDetail{
		D1: r.Intn(6) + 1,
		D2: r.Intn(6) + 1,
		D3: r.Intn(6) + 1,
	}
	Logger.Println(detail)
	return detail
}

// GetGameID returns the game ID.
func (d *DiceGame) GetGameID() GameType {
	return Dice
}
