package service_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/app/service"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestRunInn(t *testing.T) {
	s := service.GetGameResultInstance()
	int, err := s.GetLatestRunInn(1)
	assert.NoError(t, err)
	assert.NotNil(t, int)
	assert.NotEqual(t, -1, int)
}
