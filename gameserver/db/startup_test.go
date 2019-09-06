package db_test

import (
	"testing"

	"github.com/Edwardz43/mygame/gameserver/db"
)

func TestConnect(t *testing.T) {
	//TODO
	d := db.Connect()
	if err := d.Ping(); err != nil {
		t.Errorf("an error '%s' was not expected when opening db connection", err)
	}
}

func TestConnectGorm(t *testing.T) {
	//TODO
	d := db.ConnectGorm()
	if err := d.DB().Ping(); err != nil {
		t.Errorf("an error '%s' was not expected when opening db connection", err)
	}
}
