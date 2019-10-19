package repository

import (
	"github.com/jinzhu/gorm"
)

// BetResultRepository ...
type BetResultRepository struct {
	db *gorm.DB
}

// GetBetResultInstance ...
func GetBetResultInstance(db *gorm.DB) *BetResultRepository {
	return &BetResultRepository{db: db}
}
