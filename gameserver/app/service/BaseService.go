package service

import (
	"database/sql"

	"github.com/Edwardz43/mygame/gameserver/lib/log"

	"github.com/sirupsen/logrus"
)

var (
	dbConn *sql.DB
	Logger *logrus.Logger
)

func init() {
	Logger = log.Create("service")
}
