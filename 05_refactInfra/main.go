package main

import (
	"github.com/krlspj/banking-hex-arch/05_refactInfra/app"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/logger"
)

func main() {
	log := logger.NewLogger()
	logger.Logg.Info("starting app")
	log.Error("starting app")
	a := 5
	log.Infof("hello %v", a)
	logger.Info("startig the application")
	app.Start()
}
