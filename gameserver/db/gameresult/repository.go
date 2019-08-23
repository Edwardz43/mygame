package gameresult

import (
	"context"

	gameserver "github.com/Edwardz43/mygame/gameserver/app"
	"github.com/Edwardz43/mygame/gameserver/db/models"
)

// Repository represent the author's repository contract
type Repository interface {
	AddNewOne(*gameserver.GameResult) (int64, error)
	GetByBetNo(ctx context.Context, betNo int64) (*models.GameResult, error)
}
