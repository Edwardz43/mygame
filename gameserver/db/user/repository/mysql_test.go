package repository_test

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/Edwardz43/mygame/gameserver/db/user/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Updated_At", "Created_At"}).
		AddRow(1, "Ed", time.Now(), time.Now())

	query := "SELECT \\* FROM Users WHERE ID=\\?"

	userID := int64(1)

	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(userID).
		WillReturnRows(rows)

	a := repository.NewMysqlUserRepository(db)
	aUser, err := a.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, aUser)
}

func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Updated_At", "Created_At"}).
		AddRow(1, "Ed", time.Now(), time.Now())

	query := "SELECT \\* FROM Users WHERE Name=\\?"

	userName := "Ed"
	mock.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(userName).
		WillReturnRows(rows)

	a := repository.NewMysqlUserRepository(db)
	aUser, err := a.GetByName(context.TODO(), userName)
	assert.NoError(t, err)
	assert.NotNil(t, aUser)
}

func TestCreateNewOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO Users").ExpectExec().WithArgs("test").WillReturnResult(sqlmock.NewResult(1, 1))

	a := repository.NewMysqlUserRepository(db)

	var n int64

	if n, err = a.CreateNewOne(context.TODO(), "test"); err != nil {
		t.Errorf("an error '%s' was not expected when create a new user", err)
	}
	assert.NotZero(t, n)
	assert.NotEqual(t, -1, n)
}
