package repository

import (
	"context"
	"database/sql"
	"time"

	gameserver "github.com/Edwardz43/mygame/gameserver/app"
	"github.com/sirupsen/logrus"

	"github.com/Edwardz43/mygame/gameserver/db/gameresult"
	"github.com/Edwardz43/mygame/gameserver/db/models"
)

type mysqlGameResultRepo struct {
	DB *sql.DB
}

func (m *mysqlGameResultRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.GameResult, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.GameResult{}

	err = row.Scan(
		&a.ID,
		&a.GameID,
		&a.Detail,
		&a.CreatedAt,
		&a.ModTimes,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlGameResultRepo) createOne(ctx context.Context, query string, args ...interface{}) (int64, error) {
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

// NewMysqlGameResultRepository will create an implementation of user.Repository
func NewMysqlGameResultRepository(db *sql.DB) gameresult.Repository {
	return &mysqlGameResultRepo{
		DB: db,
	}
}

func (m *mysqlGameResultRepo) AddNewOne(result *gameserver.GameResult) (int64, error) {
	query := "INSERT INTO GameResult (GameID, Detail, Created_At, ModTimes) VALUES (?, ?, ?, ?)"
	return m.createOne(context.TODO(), query, result.GameType, result.GameDetail, time.Now(), 0)
}

func (m *mysqlGameResultRepo) GetByBetNo(ctx context.Context, betNo int64) (*models.GameResult, error) {
	//
	return nil, nil
}
