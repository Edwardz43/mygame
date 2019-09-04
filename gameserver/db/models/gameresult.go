package models

import "github.com/jinzhu/gorm"

// GameResult represent the gameresult model
// type GameResult struct {
// 	ID        int64  `json:"id"`
// 	GameID    int8   `json:"game_id"`
// 	Run       int64  `json:"run"`
// 	Inn       int    `json:"inn"`
// 	Detail    string `json:"detail"`
// 	CreatedAt string `json:"created_at"`
// 	ModTimes  int8   `json:"mod_times"`
// }

// GameResult represent the gameresult model
type GameResult struct {
	gorm.Model
	GameID   uint8
	Run      uint64
	Inn      uint16
	Detail   string
	ModTimes uint8
}
