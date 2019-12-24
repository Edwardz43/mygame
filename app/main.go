package main

import (
	"log"
	"net/http"

	"github.com/Edwardz43/mygame/app/config"
	"github.com/Edwardz43/mygame/app/net"

	_ "net/http/pprof"
)

func main() {

	if config.GetPprofEnable() {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	go func() {
		net.Startup()
	}()

	select {}
}
