package repository

import (
	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/jinzhu/gorm"
)

// GameResultRepository ...
type GameResultRepository struct {
	db *gorm.DB
}

// AddNewOne add
func (repo GameResultRepository) AddNewOne(gameType int8, run int64, inn int, detail string, modID int) (int64, error) {
	var gr models.GameResult

	d := repo.db.
		Create(&models.GameResult{
			GameID:   1,
			Run:      20190831,
			Inn:      1,
			Detail:   "d1:1, d2:2, d3:3",
			ModTimes: 0,
		}).
		Scan(&gr)

	if d.Error != nil {
		return 0, d.Error
	}

	return int64(gr.ID), nil
}

// GetOne ...
func (repo GameResultRepository) GetOne(gameType int8, run int64, inn int) (*models.GameResult, error) {
	var gr models.GameResult

	repo.db.First(&gr, "run = ?", 20190831)
	return nil, nil
}

// GameResult ...
// AddNewOne(gameType int8, run int64, inn int, detail string, modID int) (int64, error)
// 	GetOne(gameType int8, run int64, inn int) (*models.GameResult, error)
// 	GetByRun(gameType int8, runStart int64, runEnd int64) ([]*models.GameResult, error)
// 	GetLatestRunInn(gameType int8) (int, error)
