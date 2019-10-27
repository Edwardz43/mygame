package service_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/app/gamelogic"
	"github.com/Edwardz43/mygame/gameserver/app/service"
)

func TestSettlement(t *testing.T) {
	service := service.SettlementService{}

	diceDetails := gamelogic.DiceGameDetail{
		D1: 1, D2: 2, D3: 3,
	}

	result := gamelogic.GameResult{
		Run:        20191026,
		Inn:        1,
		GameType:   1,
		GameDetail: diceDetails,
	}

	service.Settlement(&result)
}
