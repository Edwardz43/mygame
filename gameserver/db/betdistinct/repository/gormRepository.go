package repository

import (
	"github.com/jinzhu/gorm"
)

// BetDistinctRepository ...
type BetDistinctRepository struct {
	db *gorm.DB
}

// GetBetDistinctInstance ...
func GetBetDistinctInstance(db *gorm.DB) *BetDistinctRepository {
	return &BetDistinctRepository{db: db}
}
