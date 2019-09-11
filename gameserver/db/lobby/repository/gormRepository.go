package repository

import (
	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/jinzhu/gorm"
)

// LobbyRepository ...
type LobbyRepository struct {
	db *gorm.DB
}

// GetLobbyInstance ...
func GetLobbyInstance(db *gorm.DB) *LobbyRepository {
	return &LobbyRepository{db: db}
}

// GetLatest ..
func (repo LobbyRepository) GetLatest(gameID int) (int64, int, int8, error) {
	var lobby models.Lobby
	d := repo.db.First(&lobby, "game_id = ?", gameID)
	return lobby.Run, lobby.Inn, lobby.Status, d.Error
}

// Update ..
func (repo LobbyRepository) Update(gameID int, run int64, inn int, status int) error {
	var lobby models.Lobby
	d := repo.db.First(&lobby, "game_id = ?", gameID)

	d.Model(lobby).Updates(models.Lobby{GameID: int8(gameID), Run: run, Inn: inn, Status: int8(status)})

	return d.Error
}
