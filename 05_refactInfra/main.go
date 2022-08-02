package main

import (
	"github.com/krlspj/banking-hex-arch/05_refactInfra/app"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/logger"
)

func main() {
	logger.Info("starting the application")
	app.Start()
}
