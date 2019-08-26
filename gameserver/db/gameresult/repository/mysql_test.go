package repository_test

import (
	"encoding/json"
	"fmt"
	"strconv"
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

	rows := sqlmock.NewRows([]string{"ID", "GameID", "Run", "Detail", "Created_At", "ModTimes"}).
		AddRow(1, 1, 201908260001, "", time.Now(), 0)

	query := "SELECT (.+) FROM GameResult WHERE GameID=\\? AND Run=\\?;"

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(1, 201908260001).
		WillReturnRows(rows)

	a := repository.NewMysqlGameResultRepository(db)
	result, err := a.GetOne(1, 201908260001)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetByRun(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "GameID", "Run", "Detail", "Created_At", "ModTimes"}).
		AddRow(1, 1, 201908260001, "", time.Now(), 0).
		AddRow(1, 1, 201908260002, "", time.Now(), 0)

	query := "SELECT (.+) FROM GameResult WHERE GameID=\\? AND Run BETWEEN \\? AND \\?;"

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(1, 201908260001, 201908260002).
		WillReturnRows(rows)

	a := repository.NewMysqlGameResultRepository(db)
	result, err := a.GetByRun(1, 201908260001, 201908260002)
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

	run, _ := strconv.Atoi(time.Now().Format("20060102") + fmt.Sprintf("%04d", 1))

	mock.ExpectPrepare("INSERT INTO GameResult").
		ExpectExec().
		WithArgs(1, int64(run), string(r), 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	a := repository.NewMysqlGameResultRepository(db)

	var n int64

	gr := gameserver.GameResult{
		Run:        1,
		GameType:   gameserver.Dice,
		GameDetail: detail,
	}

	r, err = json.Marshal(gr.GameDetail)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling game detail", err)
	}

	run, _ = strconv.Atoi(time.Now().Format("20060102") + fmt.Sprintf("%04d", gr.Run))

	if n, err = a.AddNewOne(int8(gr.GameType), int64(run), string(r), 0); err != nil {
		t.Errorf("an error '%s' was not expected when add a new game result", err)
	}
	assert.NotZero(t, n)
	assert.NotEqual(t, -1, n)
}
