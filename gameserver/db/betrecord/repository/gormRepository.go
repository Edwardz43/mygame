package repository

import (
	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/jinzhu/gorm"
)

// BetRecordRepository ...
type BetRecordRepository struct {
	db *gorm.DB
}

// GetBetRecordInstance ...
func GetBetRecordInstance(db *gorm.DB) *BetRecordRepository {
	return &BetRecordRepository{db: db}
}

func (repo BetRecordRepository) create(betRecord *models.BetRecord) *gorm.DB {
	return repo.db.Create(betRecord)
}

// CreateOne ...
func (repo BetRecordRepository) CreateOne(gameID int8, run int64, inn int, memberID int, distinctID int, amount int) (int, error) {
	var gr models.BetRecord

	d := repo.create(&models.BetRecord{
		GameID:     gameID,
		Run:        run,
		Inn:        inn,
		MemberID:   memberID,
		DistinctID: distinctID,
		Amount:     amount,
	}).Scan(&gr)

	if d.Error != nil {
		return -1, d.Error
	}
	return 1, nil
}
