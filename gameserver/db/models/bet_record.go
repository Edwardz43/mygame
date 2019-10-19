package models

import "github.com/jinzhu/gorm"

// BetRecord ...
type BetRecord struct {
	gorm.Model
	GameID     int8  `gorm:"game_id" json:"game_id"`
	Run        int64 `gorm:"run" json:"run"`
	Inn        int   `gorm:"inn" json:"inn"`
	MemberID   int   `gorm:"member_id" json:"member_id"`
	DistinctID int   `gorm:"distinct" json:"distinct"`
	Amount     int   `gorm:"amount" json:"amount"`
}
