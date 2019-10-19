package service

import (
	"github.com/Edwardz43/mygame/gameserver/db"
	"github.com/Edwardz43/mygame/gameserver/db/betrecord"
	"github.com/Edwardz43/mygame/gameserver/db/betrecord/repository"
)

// BetRecordService ...
type BetRecordService struct {
	Repo betrecord.Repository
}

// GetBetRecordInstance ...
func GetBetRecordInstance() *BetRecordService {
	return &BetRecordService{
		// Repo: repository.NewMysqlBetRecordRepository(db.Connect()),
		Repo: repository.GetBetRecordInstance(db.ConnectGorm()),
	}
}

// AddNewOne adds a new bet record.
func (service *BetRecordService) AddNewOne(gameID int8, run int64, inn int, memberID int, distinctID int, amount int) (message string, err error) {
	logger.Printf("parameters [%d][%d][%d][%d][%d][%d]", gameID, run, inn, memberID, distinctID, amount)
	n, err := service.Repo.CreateOne(gameID, run, inn, memberID, distinctID, amount)

	if err != nil {
		logger.Printf("BetRecordService, err : [%v]", err)
		return "err", nil
	}

	if n > 0 {
		logger.Printf("BetRecordService, AddNewOne : OK")
		return "ok", nil
	}

	logger.Printf("BetRecordService, AddNewOne : fail")
	return "fail", nil
}
