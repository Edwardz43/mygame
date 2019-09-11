package repository_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/db/gameresult/repository"
	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestGetGameResultC(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create
	db.Create(&models.GameResult{
		GameID:   1,
		Run:      20190831,
		Inn:      1,
		Detail:   "d1:1, d2:2, d3:3",
		ModTimes: 0,
	})

	var gr models.GameResult
	db.Debug().First(&gr, "run = ?", 20190831)

	assert.NotNil(t, gr)
}

func TestAddNewOneShouldReturnGameResultID(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repository.GetGameResultInstance(db)

	i, err := repo.AddNewOne(1, 20190906, 1, "{d1:3, d2:6, d3:1}", 0)

	if err != nil {
		panic(err)
	}

	assert.NotEqual(t, i, -1)
}

func TestAddNewOneShouldReturnError(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repository.GetGameResultInstance(db)

	i, err := repo.AddNewOne(1, 20190906, 1, "{d1:3, d2:6, d3:1}", 0)

	if err != nil {
		panic(err)
	}

	assert.NotEqual(t, i, -1)
}
