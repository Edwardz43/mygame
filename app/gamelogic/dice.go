package gamelogic

import (
	"math/rand"
	"time"
)

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
	logger.Println("Start Game Dice")
}

// NewGame ...
func (d *DiceGame) NewGame() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	detail := DiceGameDetail{
		D1: r.Intn(6) + 1,
		D2: r.Intn(6) + 1,
		D3: r.Intn(6) + 1,
	}
	// logger.Println(detail)
	return detail
}

// GetGameID returns the game ID.
func (d *DiceGame) GetGameID() GameType {
	return Dice
}
