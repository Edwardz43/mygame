package service

import (
	"log"

	"github.com/Edwardz43/mygame/gameserver/app/gamelogic"
	"github.com/Edwardz43/mygame/gameserver/db/betdistinct"
	"github.com/Edwardz43/mygame/gameserver/db/betresult"
)

// SettlementService ...
type SettlementService struct {
	BetDistinctRepo betdistinct.Repository
	BetResultRepo   betresult.Repository
}

// Settlement ...
func (service *SettlementService) Settlement(result *gamelogic.GameResult) {

	switch result.GameType {
	case gamelogic.Dice:
		detail := result.GameDetail.(gamelogic.DiceGameDetail)
		s := &diceSettlement{}
		s.settlement(&detail)
		break
	default:
		//TODO
		break
	}

}

type settlementLogic interface {
	settlement()
}

// updateResult ...
func updateResult(logic settlementLogic) {
	logic.settlement()
}

type diceSettlement struct{}

func (s *diceSettlement) settlement(detail *gamelogic.DiceGameDetail) {
	d1 := detail.D1
	d2 := detail.D2
	d3 := detail.D3
	log.Printf("d1:%d, d2:%d, d3:%d\n", d1, d2, d3)

	//大小
	big := (d1 + d2 + d3) > 10

	//單雙
	odd := (d1+d2+d3)%2 == 0

	log.Printf("big:%v, odd:%v\n", big, odd)
}
