package gameresult

import (
	"context"

	"github.com/Edwardz43/mygame/gameserver/db/models"
)

// Repository represent the author's repository contract
type Repository interface {
	AddNewOne(gameType int8, run int64, detail string, modID int) (int64, error)
	GetByBetNo(ctx context.Context, betNo int64) (*models.GameResult, error)
}
