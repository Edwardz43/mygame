package repository_test

import (
	"testing"

	"github.com/Edwardz43/mygame/app/db/member/repository"
	"github.com/Edwardz43/mygame/gameserver/db/models"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

func TestLoginWithNameShouldReturnSuccess(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Member{})

	repo := repository.GetMemberInstance(db)

	id, err := repo.GetOne("test001")

	assert.Empty(t, err)
	assert.NotEqual(t, 0, id)
}

func TestLoginWithEmailShouldReturnSuccess(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.GetMemberInstance(db)

	id, err := repo.GetOne("test001@com.tw")

	assert.Empty(t, err)
	assert.NotEqual(t, 0, id)
}

func TestLoginShouldReturnErr(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.GetMemberInstance(db)

	_, err = repo.GetOne("xxx")

	assert.NotNil(t, err)
}

func TestRegisterShouldReturnSuccess(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.GetMemberInstance(db)

	ok, err := repo.Create("test006", "test006@test.com", "8888")

	assert.Nil(t, err)
	assert.True(t, ok)
}
