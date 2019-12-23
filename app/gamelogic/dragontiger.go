package gamelogic

import (
	"math/rand"
	"time"
)

// DragonTigerGameDetail ...
type DragonTigerGameDetail struct {
	BankerCard string `json:"b_card"`
	PlayerCard string `json:"p_card"`
}

// DragonTigerGame ...
type DragonTigerGame struct{}

// StartGame ...
func (d *DragonTigerGame) StartGame() {
	logger.Println("Start Game")
}

// NewGame ...
func (d *DragonTigerGame) NewGame() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	detail := DragonTigerGameDetail{
		BankerCard: Poker[uint8(r.Intn(52))],
		PlayerCard: Poker[uint8(r.Intn(52))],
	}
	// logger.Println(detail)
	return detail
}

// GetGameID returns the game ID.
func (d *DragonTigerGame) GetGameID() GameType {
	return DragonTiger
}
