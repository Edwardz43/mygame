package service

import (
	"database/sql"

	"github.com/Edwardz43/mygame/gameserver/db"
	"github.com/Edwardz43/mygame/gameserver/lib/log"
	"github.com/jinzhu/gorm"
)

var (
	dbConn     *sql.DB
	dbGormConn *gorm.DB
	logger     *log.Logger
)

func init() {
	logger = log.Create("service")
	dbGormConn = db.ConnectGorm()
}
