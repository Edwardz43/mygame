package service

import (
	"github.com/Edwardz43/mygame/app/lib/log"
	"github.com/Edwardz43/mygame/app/db"
	"database/sql"
	
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
