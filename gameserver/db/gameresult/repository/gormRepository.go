package repository

import (
	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/jinzhu/gorm"
)

// GameResultRepository ...
type GameResultRepository struct {
	db *gorm.DB
}

// GetGameResultInstance ...
func GetGameResultInstance(db *gorm.DB) *GameResultRepository {
	return &GameResultRepository{db: db}
}

func (repo GameResultRepository) create(result *models.GameResult) *gorm.DB {
	return repo.db.Create(result)
}

// AddNewOne add
func (repo GameResultRepository) AddNewOne(gameType int8, run int64, inn int, detail string, modID int8) (int64, error) {
	var gr models.GameResult

	d := repo.create(&models.GameResult{
		GameID:   gameType,
		Run:      run,
		Inn:      inn,
		Detail:   detail,
		ModTimes: modID,
	}).Scan(&gr)

	if d.Error != nil {
		return -1, d.Error
	}
	return int64(gr.ID), nil
}

// GetOne ...
func (repo GameResultRepository) GetOne(gameType int8, run int64, inn int) (*models.GameResult, error) {
	var gr models.GameResult

	repo.db.First(&gr, "run = ?", 20190831)
	return nil, nil
}

// GetByRun ...
func (repo GameResultRepository) GetByRun(gameType int8, runStart int64, runEnd int64) ([]*models.GameResult, error) {
	return nil, nil
}
