package repository_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/db/lobby/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestShouldReturnSuccess(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.GetLobbyInstance(db)

	run, inn, status, err := repo.GetLatest(1)

	assert.Empty(t, err)
	assert.NotEqual(t, 0, run)
	assert.NotEqual(t, 0, inn)
	assert.NotEqual(t, 0, status)
}

func TestGetLatestShouldReturnErr(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.GetLobbyInstance(db)

	_, _, _, err = repo.GetLatest(2)

	assert.NotNil(t, err)
}

func TestUpdateShouldReturnSuccess(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.GetLobbyInstance(db)

	run, inn, status, _ := repo.GetLatest(1)

	err = repo.Update(1, 20190911, 2, 1)
	assert.Empty(t, err)

	err = repo.Update(1, run, inn, int(status))
	assert.Empty(t, err)
}
