package repository_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/app/repository"
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

func TestGetGameResult(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&repository.GameResult{})

	// Create
	db.Create(&repository.GameResult{
		GameID:   1,
		Run:      20190831,
		Inn:      1,
		Detail:   "d1:1, d2:2, d3:3",
		ModTimes: 0,
	})

	var gr repository.GameResult
	db.First(&gr, "run = ?", 20190831)

	db.Model(&gr).Update("inn", 2)

	db.Delete(&gr)
	assert.NotNil(t, gr)
}

// func TestProduct(t *testing.T) {
// 	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	// Migrate the schema
// 	db.AutoMigrate(&repository.Product{})

// 	// Create
// 	db.Create(&repository.Product{Code: "L1212", Price: 1000})

// 	// Read
// 	var product repository.Product
// 	db.First(&product, 1)                   // find product with id 1
// 	db.First(&product, "code = ?", "L1212") // find product with code l1212

// 	// Update - update product's price to 2000
// 	db.Model(&product).Update("Price", 2000)

// 	// Delete - delete product
// 	db.Delete(&product)
// }
