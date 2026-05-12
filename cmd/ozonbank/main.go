package main

import (
	"OzonBank/internal/app"
	"OzonBank/internal/utils"
)

func main() {
	// Initialize logger
	logger := utils.NewLogger()

	// Load configuration
	config := app.LoadConfig()

	// Create and run application
	application := app.NewApplication(config, logger)
	application.Run()
}
