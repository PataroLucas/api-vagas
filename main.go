package main

import (
	"github.com/PataroLucas/api-vagas/config"
	"github.com/PataroLucas/api-vagas/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	error := config.Init()
	if error != nil {
		logger.Errorf("config initialization error: %v", error)
		return
	}

	router.Initialize()
}
