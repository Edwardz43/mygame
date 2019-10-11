package service

import (
	"github.com/Edwardz43/mygame/gameserver/db"
	"github.com/Edwardz43/mygame/gameserver/db/gameresult"
	"github.com/Edwardz43/mygame/gameserver/db/gameresult/repository"
)

// GameResultService ...
type GameResultService struct {
	Repo gameresult.Repository
}

// GetGameResultInstance ...
func GetGameResultInstance() *GameResultService {
	return &GameResultService{
		// Repo: repository.NewMysqlGameResultRepository(db.Connect()),
		Repo: repository.GetGameResultInstance(db.ConnectGorm()),
	}
}

// AddNewOne add a new result.
func (service *GameResultService) AddNewOne(gameType int8, run int64, inn int, detail string, modID int8) (message string, err error) {
	// defer service.dbConn.Close()
	Logger.Printf("parameters [%d][%d][%d][%s][%d]", gameType, run, inn, detail, modID)
	n, err := service.Repo.AddNewOne(gameType, run, inn, detail, modID)

	if err != nil {
		Logger.Println(err)
		return "err", nil
	}

	if n > 0 {
		return "ok", nil
	}

	return "fail", nil
}
