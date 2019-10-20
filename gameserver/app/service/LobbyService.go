package service

import (
	"github.com/Edwardz43/mygame/gameserver/db"
	"github.com/Edwardz43/mygame/gameserver/db/lobby"
	"github.com/Edwardz43/mygame/gameserver/db/lobby/repository"
)

// LobbyService ...
type LobbyService struct {
	Repo lobby.Repository
}

// GetLobbyInstance returns instance of lobby service.
func GetLobbyInstance() *LobbyService {
	return &LobbyService{
		// Repo: repository.NewMysqlGameResultRepository(db.Connect()),
		Repo: repository.GetLobbyInstance(db.ConnectGorm()),
	}
}

// GetLatest returns the latest game info with specific game ID.
func (service *LobbyService) GetLatest(gameID int) (int64, int, int8, int8, error) {
	logger.Printf("parameters [%d]", gameID)
	return service.Repo.GetLatest(gameID)
}

// Update updates the info.
func (service *LobbyService) Update(gameID int, run int64, inn int, status int) error {
	logger.Printf("parameters [%d][%d][%d][%d]", gameID, run, inn, status)
	if status == 2 {

	}
	return service.Repo.Update(gameID, run, inn, status)
}

// Countdown updates the countdown.
func (service *LobbyService) Countdown(gameID int, countdown int8) error {
	logger.Printf("parameters [%d][%d]", gameID, countdown)
	return service.Repo.Countdown(gameID, countdown)
}
