package repository

import (
	"github.com/jinzhu/gorm"
)

type GameResult struct {
	gorm.Model
	GameID   uint8
	Run      uint64
	Inn      uint16
	Detail   string
	ModTimes uint8
}

func test() {
	db, err := gorm.Open("mysql", "root:root@/MyGame")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	// db.AutoMigrate(&GameResult{})

	// Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var gr GameResult
	db.First(&gr, 1)                   // find product with id 1
	db.First(&gr, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&gr).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&gr)
}
