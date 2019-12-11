package gamelogic_test

import (
	"github.com/Edwardz43/mygame/app/gamelogic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := gamelogic.DiceGame{}
	gameResult := g.NewGame()
	assert.NotNil(t, gameResult)
	// assert.NotNil(t, gameResult.GameType)
	// assert.NotNil(t, gameResult.GameDetail)
}
