package main

import (
	"github.com/krlspj/banking-hex-arch/06_wru_t/app"
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/logger"
)

func main() {
	logger.Info("starting the application")
	app.Start()
}
