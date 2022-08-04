package main

import (
	"github.com/krlspj/banking-hex-arch/07_security/app"
	"github.com/krlspj/banking-hex-arch/07_security/internal/logger"
)

func main() {
	logger.Info("starting the application")
	app.Start()
}
