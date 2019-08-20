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
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlUserRepo) createOne(ctx context.Context, query string, args ...interface{}) (int64, error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return -1, err
	}
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		logrus.Error(err)
		return -1, err
	}
	a, err := result.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return -1, err
	}
	return a, nil
}

func (m *mysqlUserRepo) CreateNewOne(ctx context.Context, name string) (int64, error) {
	query := "INSERT INTO Users (Name) VALUES (?)"
	return m.createOne(ctx, query, name)
}

func (m *mysqlUserRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT * FROM Users WHERE ID=?`
	return m.getOne(ctx, query, id)
}

func (m *mysqlUserRepo) GetByName(ctx context.Context, name string) (*models.User, error) {
	query := `SELECT * FROM Users WHERE Name=?`
	return m.getOne(ctx, query, name)
}
