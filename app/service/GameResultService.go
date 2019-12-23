package service

import (
	"github.com/Edwardz43/mygame/app/db/gameresult"
	"github.com/Edwardz43/mygame/app/db/gameresult/repository"
)

// GameResultService ...
type GameResultService struct {
	Repo gameresult.Repository
}

// GetGameResultInstance ...
func GetGameResultInstance() *GameResultService {
	return &GameResultService{
		// Repo: repository.NewMysqlGameResultRepository(db.Connect()),
		Repo: repository.GetGameResultInstance(dbGormConn),
	}
}

// AddNewOne add a new result.
func (service *GameResultService) AddNewOne(gameType int8, run int64, inn int, detail string, modID int8) (message string, err error) {
	// logger.Printf("parameters [%d][%d][%d][%s][%d]", gameType, run, inn, detail, modID)
	n, err := service.Repo.AddNewOne(gameType, run, inn, detail, modID)

	if err != nil {
		logger.Println(err.Error())
		return "err", nil
	}

	if n > 0 {
		return "ok", nil
	}

	return "fail", nil
}

// GetLatest get the latest game result.
func (service *GameResultService) GetLatest(gameType int8, run int64, inn int) (result string, err error) {
	// logger.Printf("parameters [%d][%d][%d]", gameType, run, inn)

	model, err := service.Repo.GetOne(gameType, run, inn)

	if err != nil {
		logger.Println(err.Error())
		return "err", err
	}

	return model.Detail, nil
}
