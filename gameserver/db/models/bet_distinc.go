package models

import "github.com/jinzhu/gorm"

// Distinct ...
type Distinct struct {
	gorm.Model
	GameID   int8   `gorm:"game_id" json:"game_id"`
	Distinct string `gorm:"distinct" json:"distinct"`
	WinFlag  bool   `gorm:"win_flag" json:"win_flag"`
}
