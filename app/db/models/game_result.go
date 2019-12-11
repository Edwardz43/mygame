package models

import (
	"github.com/jinzhu/gorm"
)

// GameResult represent the gameresult model
type GameResult struct {
	gorm.Model
	// ID        int64      `gorm:"id" json:"id"`
	GameID int8   `gorm:"game_id" json:"game_id"`
	Run    int64  `gorm:"run" json:"run"`
	Inn    int    `gorm:"inn" json:"inn"`
	Detail string `gorm:"detail" json:"detail"`
	// CreatedAt *time.Time `gorm:"created_at" json:"created_at"`
	// UpdatedAt *time.Time `gorm:"updated_at" json:"updated_at"`
	ModTimes int8 `gorm:"mod_times" json:"mod_times"`
}

// TableName returns the custom table name.
func (GameResult) TableName() string {
	return "game_result"
}
