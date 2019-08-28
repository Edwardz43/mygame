package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/Edwardz43/mygame/gameserver/db/gameresult"
	"github.com/Edwardz43/mygame/gameserver/db/models"
)

type mysqlGameResultRepo struct {
	DB *sql.DB
}

func (m *mysqlGameResultRepo) getOne(query string, args ...interface{}) (*models.GameResult, error) {

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRow(args...)
	a := &models.GameResult{}

	err = row.Scan(
		&a.ID,
		&a.GameID,
		&a.Run,
		&a.Inn,
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

func (m *mysqlGameResultRepo) getMany(query string, args ...interface{}) ([]*models.GameResult, error) {
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	list := make([]*models.GameResult, 0)

	defer rows.Close()
	for rows.Next() {
		a := &models.GameResult{}
		err = rows.Scan(
			&a.ID,
			&a.GameID,
			&a.Run,
			&a.Inn,
			&a.Detail,
			&a.CreatedAt,
			&a.ModTimes,
		)
		if err != nil {
			// handle this error
			panic(err)
		}
		list = append(list, a)
	}

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return list, nil
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

func (m *mysqlGameResultRepo) AddNewOne(gameType int8, run int64, inn int, detail string, modID int) (int64, error) {
	query := "INSERT INTO GameResult (GameID, Run, Inn, Detail, ModTimes) VALUES (?, ?, ?, ?, ?);"
	return m.createOne(context.TODO(), query, int8(gameType), run, inn, detail, modID)
}

func (m *mysqlGameResultRepo) GetOne(gameType int8, run int64, inn int) (*models.GameResult, error) {
	query := "SELECT * FROM GameResult WHERE GameID=? AND Run=? AND Inn=?;"
	return m.getOne(query, gameType, run, inn)
}

func (m *mysqlGameResultRepo) GetByRun(gameType int8, runStart int64, runEnd int64) ([]*models.GameResult, error) {
	query := "SELECT * FROM GameResult WHERE GameID=? AND Run BETWEEN ? AND ?;"
	return m.getMany(query, gameType, runStart, runEnd)
}

func (m *mysqlGameResultRepo) GetLatestRunInn(gameType int8) (int, error) {
	query := "SELECT * FROM GameResult WHERE GameID=? ORDER BY ID DESC LIMIT 1;"
	gr, err := m.getOne(query, gameType)
	if err != nil {
		return -1, err
	}
	return gr.Inn, nil
}
