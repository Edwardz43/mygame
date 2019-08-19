package repository

import (
	"context"
	"database/sql"

	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/Edwardz43/mygame/gameserver/db/user"
	"github.com/sirupsen/logrus"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

// NewMysqlUserRepository will create an implementation of user.Repository
func NewMysqlUserRepository(db *sql.DB) user.Repository {
	return &mysqlUserRepo{
		DB: db,
	}
}

func (m *mysqlUserRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.User, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.User{}

	err = row.Scan(
		&a.ID,
		&a.Name,
		// &a.CreatedAt,
		// &a.UpdatedAt,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlUserRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT id, name FROM Users WHERE id=?`
	return m.getOne(ctx, query, id)
}
