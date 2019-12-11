package repository

import (
	"strconv"
	"time"

	"github.com/Edwardz43/mygame/app/db/models"
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

// Create ..
func (repo LobbyRepository) Create(gameID int8) (bool, error) {
	run, _ := strconv.Atoi(time.Now().Format("20060102"))

	lobby := &models.Lobby{
		GameID: gameID,
		Run:    int64(run),
		Inn:    0,
		Status: 1,
	}

	d := repo.db.Create(lobby)
	if d.Error != nil {
		return false, d.Error
	}

	return true, nil
}

// GetLatest ..
func (repo LobbyRepository) GetLatest(gameID int) (int64, int, int8, int8, error) {
	var lobby models.Lobby
	d := repo.db.First(&lobby, "game_id = ?", gameID)
	return lobby.Run, lobby.Inn, lobby.Status, lobby.Countdown, d.Error
}

// Update ..
func (repo LobbyRepository) Update(gameID int, run int64, inn int, status int) error {
	var lobby models.Lobby
	d := repo.db.First(&lobby, "game_id = ?", gameID)

	d.Model(lobby).Updates(map[string]interface{}{"game_id": int8(gameID), "run": run, "inn": inn, "status": int8(status)})

	return d.Error
}

// Countdown ...
func (repo LobbyRepository) Countdown(gameID int, countdown int8) error {
	var lobby models.Lobby
	d := repo.db.First(&lobby, "game_id = ?", gameID)

	d.Model(lobby).Update("countdown", countdown)

	return d.Error
}
