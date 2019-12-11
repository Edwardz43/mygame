package models

import "github.com/jinzhu/gorm"

// BetResult ...
type BetResult struct {
	gorm.Model
	BetRecordID int8   `gorm:"bet_record_id" json:"bet_record_id"`
	WinFlag     bool   `gorm:"win_flag" json:"win_flag"`
	WinDistinct string `gorm:"win_distinct" json:"win_distinct"`
	Distinct    string `gorm:"distinct" json:"distinct"`
	BetAmount   int    `gorm:"bet_amount" json:"bet_amount"`
	BetWin      int    `gorm:"bet_win" json:"bet_win"`
}
