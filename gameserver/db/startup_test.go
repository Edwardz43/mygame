package db_test

import (
	"log"
	"testing"

	gameserver "github.com/Edwardz43/mygame/gameserver/app"
	"github.com/Edwardz43/mygame/gameserver/db/gameresult/repository"

	"github.com/Edwardz43/mygame/gameserver/db"
)

func TestConnect(t *testing.T) {
	//TODO
	d := db.Connect()
	if err := d.Ping(); err != nil {
		t.Errorf("an error '%s' was not expected when opening db connection", err)
	}

	gr := repository.NewMysqlGameResultRepository(d)

	a := gameserver.GameResult{
		GameType:   gameserver.Dice,
		Run:        1,
		GameDetail: "{d1:1, d2:2, d3:3}",
	}
	n, err := gr.AddNewOne(&a)
	if err != nil {
		t.Errorf("an error '%s' was not expected when add a new game result", err)
	}
	log.Println(n)

}
