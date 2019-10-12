package gameserver_test

import (
	"testing"

	gameserver "github.com/Edwardz43/mygame/gameserver/app"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := gameserver.DiceGame{}
	gameResult := g.NewGame()
	assert.NotNil(t, gameResult)
	// assert.NotNil(t, gameResult.GameType)
	// assert.NotNil(t, gameResult.GameDetail)
}
