package main

import (
	"map-friend/src/interface/router"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	if os.Getenv("APP_ENV") == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	log.Info("start server port:", os.Getenv("APP_PORT"), "\nenv:", os.Getenv("APP_ENV"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), router.New())
}
