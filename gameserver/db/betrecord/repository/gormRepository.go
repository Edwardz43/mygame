package repository

import (
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
