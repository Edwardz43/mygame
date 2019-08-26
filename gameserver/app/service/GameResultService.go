package service

import (
	"database/sql"
	"log"

	"github.com/Edwardz43/mygame/gameserver/db/gameresult/repository"
)

// GameResultService ...
type GameResultService struct {
	DbConn *sql.DB
}

// AddNewOne add a new result.
func (service *GameResultService) AddNewOne(gameType int8, run int64, detail string, modID int) (message string, err error) {
	gameResult := repository.NewMysqlGameResultRepository(service.DbConn)
	// defer service.dbConn.Close()
	n, err := gameResult.AddNewOne(gameType, run, detail, modID)

	if err != nil {
		log.Println(err)
		return "err", nil
	}

	if n > 0 {
		return "ok", nil
	}

	return "fail", nil
}
