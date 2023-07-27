package main

import (
	"fmt"

	"github.com/JoaoPaulo-creator/gopportunities/config"
	"github.com/JoaoPaulo-creator/gopportunities/router"
)

func main() {
	// inicializando config
	err := config.Init()
	// matando a aplicacao, caso de ruim no Init()
	if err != nil {
		fmt.Errorf("An error occured while initializating", err)
		return
	}
	router.Initialize()
}
