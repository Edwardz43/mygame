package gamelogic

import (
	"math/rand"
	"time"
)

// DragonTigerGameDetail ...
type DragonTigerGameDetail struct {
	DragonCard string `json:"d_card"`
	TigerCard  string `json:"t_card"`
}

// DragonTigerGame ...
type DragonTigerGame struct{}

// StartGame ...
func (d *DragonTigerGame) StartGame() {
	logger.Println("Start Game DT")
}

// NewGame ...
func (d *DragonTigerGame) NewGame() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	detail := DragonTigerGameDetail{
		DragonCard: Poker[uint8(r.Intn(52))],
		TigerCard:  Poker[uint8(r.Intn(52))],
	}
	// logger.Println(detail)
	return detail
}

// GetGameID returns the game ID.
func (d *DragonTigerGame) GetGameID() GameType {
	return DragonTiger
}
