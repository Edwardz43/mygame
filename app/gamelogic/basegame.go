package gamelogic

import (
	"github.com/Edwardz43/mygame/app/lib/log"
)

// GameType ...
type GameType int8

// BetDistinct defines distinct of bet
// type BetDistinct int

var (
	logger *log.Logger
	// BetDistinctMap map[BetDistinct]string
)

const (
	Dice GameType = iota + 1
	Roulette
	DragonTiger
)

const (
	Big   string = "big"
	Small string = "small"
	Odd   string = "odd"
	Even  string = "even"
)

// GameResult ...
type GameResult struct {
	Run        int64       `json:"run"`
	Inn        int         `json:"inn"`
	GameType   GameType    `json:"game_type"`
	GameDetail interface{} `json:"game_detail"`
}

// GameBase ...
type GameBase interface {
	StartGame()
	NewGame() interface{}
	GetGameID() GameType
}

func init() {
	logger = log.Create("game_logic")
}
