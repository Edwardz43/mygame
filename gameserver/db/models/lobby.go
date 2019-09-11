package models

// Lobby ...
type Lobby struct {
	GameID int8  `gorm:"game_id" json:"game_id"`
	Run    int64 `gorm:"run" json:"run"`
	Inn    int   `gorm:"inn" json:"inn"`
	Status int8  `gorm:"status" json:"status"`
}

// TableName returns the custom table name.
func (Lobby) TableName() string {
	return "Lobby"
}
