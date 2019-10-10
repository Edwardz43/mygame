package models

import "github.com/jinzhu/gorm"

// Lobby ...
type Lobby struct {
	gorm.Model
	GameID int8  `gorm:"type:smallint;not null" json:"game_id"`
	Run    int64 `gorm:"type:bigint;not null" json:"run"`
	Inn    int   `gorm:"inn;not null" json:"inn"`
	Status int8  `gorm:"type:smallint;not null" json:"status"`
}

// TableName returns the custom table name.
func (Lobby) TableName() string {
	return "lobby"
}
