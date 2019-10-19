package gamelogic_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/app/gamelogic"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := gamelogic.DiceGame{}
	gameResult := g.NewGame()
	assert.NotNil(t, gameResult)
	// assert.NotNil(t, gameResult.GameType)
	// assert.NotNil(t, gameResult.GameDetail)
}
