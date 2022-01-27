package main

import (
	"github.com/ameydev/banking-app/app"
	"github.com/ameydev/banking-app/logger"
)

func main() {
	logger.Info("Starting the application..")
	app.Start()
}
