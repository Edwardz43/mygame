package gameserver

// GameType ...
type GameType int8
type GameStatus int8

const (
	Dice GameType = iota + 1
	Roulette
)

const (
	NewRun GameStatus = iota + 1
	Showdown
	Settlement
	Intermission
	Maintain
)

// GameResult ...
type GameResult struct {
	Run        int64       `json:"run"`
	Inn        int         `json:"inn"`
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
