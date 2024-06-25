package main

import (
	"github.com/Chrystianls/Opportunities-API.git/config"
	"github.com/Chrystianls/Opportunities-API.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
