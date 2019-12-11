package service

import (
	"github.com/Edwardz43/mygame/app/db/betrecord"
	"github.com/Edwardz43/mygame/app/db/betrecord/repository"
)

// BettingService ...
type BettingService struct {
	Repo betRecord.Repository
}

// GetBettingInstance ...
func GetBettingInstance() *BettingService {
	return &BettingService{
		// Repo: repository.NewMysqlBettingRepository(db.Connect()),
		Repo: repository.GetBetRecordInstance(dbGormConn),
	}
}

// AddNewOne adds a new bet record.
func (service *BettingService) AddNewOne(gameID int8, run int64, inn int, memberID int, distinctID int, amount int) (message string, err error) {
	logger.Printf("parameters [%d][%d][%d][%d][%d][%d]", gameID, run, inn, memberID, distinctID, amount)
	n, err := service.Repo.CreateOne(gameID, run, inn, memberID, distinctID, amount)

	if err != nil {
		logger.Printf("BettingService, err : [%v]", err)
		return "err", nil
	}

	if n > 0 {
		logger.Printf("BettingService, AddNewOne : OK")
		return "ok", nil
	}

	logger.Printf("BettingService, AddNewOne : fail")
	return "fail", nil
}
