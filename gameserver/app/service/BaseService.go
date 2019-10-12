package service

import (
	"database/sql"

	"github.com/Edwardz43/mygame/gameserver/lib/log"
)

var (
	dbConn *sql.DB
	logger *log.Logger
)

func init() {
	logger = log.Create("service")
}
