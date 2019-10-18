package models

import "github.com/jinzhu/gorm"

// BetRecord ...
type BetRecord struct {
	gorm.Model
	GameID   int8  `gorm:"game_id" json:"game_id"`
	Run      int64 `gorm:"run" json:"run"`
	Inn      int   `gorm:"inn" json:"inn"`
	Distinct int8  `gorm:"distinct" json:"distinct"`
	Amount   int64 `gorm:"amount" json:"amount"`
}
