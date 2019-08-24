package gameserver

// GameType ...
type GameType int8

const (
	Dice GameType = iota + 1
	Roulette
)

// GameResult ...
type GameResult struct {
	Run        int         `json:"run"`
	GameType   GameType    `json:"game_type"`
	GameDetail interface{} `json:"game_detail"`
}

// DiceGameDetail ...
type DiceGameDetail struct {
	D1 int `json:"d1"`
	D2 int `json:"d2"`
	D3 int `json:"d3"`
}

// GameBase ...
type GameBase interface {
	StartGame(chan *GameResult)
	NewGame() *GameResult
}
