package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Edwardz43/mygame/gameserver/config"
	"github.com/Edwardz43/mygame/gameserver/db/models"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// Connect ...
func Connect() *sql.DB {
	connection := config.GetDBConfig()
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Taipei")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return dbConn
}

// ConnectGorm ...
func ConnectGorm() *gorm.DB {
	connection := config.GetDBConfigV2()
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}

	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db.AutoMigrate(&models.GameResult{})
	db.AutoMigrate(&models.Lobby{})
	db.AutoMigrate(&models.Member{})

	return db
}
