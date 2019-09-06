package service

import (
	"log"

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
	n, err := service.Repo.AddNewOne(gameType, run, inn, detail, modID)

	if err != nil {
		log.Println(err)
		return "err", nil
	}

	if n > 0 {
		return "ok", nil
	}

	return "fail", nil
}

// GetLatestRunInn ...
func (service *GameResultService) GetLatestRunInn(gameType int8) (message int, err error) {
	n, err := service.Repo.GetLatestRunInn(gameType)

	if err != nil {
		log.Println(err)
		return -1, err
	}

	if n > 0 {
		return n, nil
	}

	return -1, nil
}
