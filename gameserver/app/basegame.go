package gameserver

// GameType ...
type GameType int8
type GameStatus int8

const (
	Dice GameType = iota + 1
	Roulette
)

const (
	NewInn GameStatus = iota + 1
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

// GameBase ...
type GameBase interface {
	StartGame()
	NewGame() interface{}
	GetGameID() GameType
}
