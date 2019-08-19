package db_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/db"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	//TODO
	d := db.Connect()
	err := d.Ping()
	assert.Empty(t, err)
}
