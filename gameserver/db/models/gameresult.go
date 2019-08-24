package models

// GameResult represent the gameresult model
type GameResult struct {
	ID        int64  `json:"id"`
	GameID    int8   `json:"game_id"`
	Run       int64  `json:"run"`
	Detail    string `json:"detail"`
	CreatedAt string `json:"created_at"`
	ModTimes  int8   `json:"mod_times"`
}
