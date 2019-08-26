package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Edwardz43/mygame/gameserver/config"
	_ "github.com/go-sql-driver/mysql"
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
