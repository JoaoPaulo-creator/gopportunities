package main

import (
	"github.com/JoaoPaulo-creator/gopportunities/config"
	"github.com/JoaoPaulo-creator/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// inicializando config
	err := config.Init()
	// matando a aplicacao, caso de ruim no Init()
	if err != nil {
		//%v serve para aceitar qualquer tipo
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
