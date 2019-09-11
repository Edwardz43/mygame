package repository_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	gameserver "github.com/Edwardz43/mygame/gameserver/app"
	"github.com/Edwardz43/mygame/gameserver/db/gameresult/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "GameID", "Run", "Inn", "Detail", "Created_At", "ModTimes"}).
		AddRow(1, 1, 20190826, 1, "", time.Now(), 0)

	query := "SELECT (.+) FROM GameResult WHERE GameID=\\? AND Run=\\? AND Inn=\\?;"

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(1, 20190826, 1).
		WillReturnRows(rows)

	a := repository.NewMysqlGameResultRepository(db)
	result, err := a.GetOne(1, 20190826, 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetByRun(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "GameID", "Run", "Inn", "Detail", "Created_At", "ModTimes"}).
		AddRow(1, 1, 20190826, 1, "", time.Now(), 0).
		AddRow(1, 1, 20190826, 2, "", time.Now(), 0)

	query := "SELECT (.+) FROM GameResult WHERE GameID=\\? AND Run BETWEEN \\? AND \\?;"

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(1, 20190826, 20190826).
		WillReturnRows(rows)

	a := repository.NewMysqlGameResultRepository(db)
	result, err := a.GetByRun(1, 20190826, 20190826)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(result), 2)
}

func TestAddNewOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	detail := gameserver.DiceGameDetail{D1: 1, D2: 2, D3: 3}

	r, err := json.Marshal(detail)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling game detail", err)
	}

	mock.ExpectPrepare("INSERT INTO GameResult").
		ExpectExec().
		WithArgs(1, 20190826, 1, string(r), 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	a := repository.NewMysqlGameResultRepository(db)

	var n int64

	gr := gameserver.GameResult{
		Run:        20190826,
		Inn:        1,
		GameType:   gameserver.Dice,
		GameDetail: detail,
	}

	r, err = json.Marshal(gr.GameDetail)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling game detail", err)
	}

	if n, err = a.AddNewOne(int8(gr.GameType), gr.Run, gr.Inn, string(r), 0); err != nil {
		t.Errorf("an error '%s' was not expected when add a new game result", err)
	}
	assert.NotZero(t, n)
	assert.NotEqual(t, -1, n)
}
