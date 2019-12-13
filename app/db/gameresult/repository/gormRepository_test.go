package repository_test

import (
	"github.com/Edwardz43/mygame/app/db/gameresult/repository"
	"github.com/Edwardz43/mygame/app/db/models"
	"testing"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestGetGameResultC(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.GameResult{})

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
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
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
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
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

func TestGetOneShouldReturnSuccess(t *testing.T) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.GetGameResultInstance(db)

	model, err := repo.GetOne(1, 20191024, 100)

	if err != nil {
		panic(err)
	}

	assert.NotNil(t, model)
	assert.NotNil(t, model.Detail)
	assert.NotEqual(t, "", model.Detail)
}
