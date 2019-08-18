package main

import (
	"log"
	"math/rand"
	"time"
)

var run = 1

type GameResult struct {
	Run int `json:"run"`
	D1  int `json:"d1"`
	D2  int `json:"d2"`
	D3  int `json:"d3"`
}

func StartGame(result chan *GameResult) {
	t := time.NewTicker(time.Second * 30)
	log.Println("Start Game")
	for {
		select {
		case <-t.C:
			log.Println("New Game")
			r := NewGame()
			// log.Println(r)
			result <- r
		default:
			//
		}
	}

}

func NewGame() *GameResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	gr := new(GameResult)
	gr.Run = run
	gr.D1 = r.Intn(6) + 1
	gr.D2 = r.Intn(6) + 1
	gr.D3 = r.Intn(6) + 1
	run++
	log.Println(gr)
	return gr
}
