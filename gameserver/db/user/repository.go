package user

import (
	"context"

	"github.com/Edwardz43/mygame/gameserver/db/models"
)

// Repository represent the author's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByName(ctx context.Context, name string) (*models.User, error)
	CreateNewOne(ctx context.Context, name string) (int64, error)
}
