package gameserver_test

import (
	"testing"

	gameserver "github.com/Edwardz43/mygame/gameserver/app"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	r := gameserver.NewGame()
	assert.NotNil(t, r.D1)
	assert.NotNil(t, r.D2)
	assert.NotNil(t, r.D3)
	assert.NotNil(t, r.Run)
}

// func TestStartGame(t *testing.T) {
// 	r := make(chan *gameserver.GameResult)

// 	go gameserver.StartGame(r)

// 	for {
// 		<-r
// 		assert.NotNil(t, r)
// 		break
// 	}
// }
