package main

import (
	"github.com/imrochamatheus/challenge-back-end-hit/config"
	"github.com/imrochamatheus/challenge-back-end-hit/router"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err.Error())
	}

	router.Initialize()
}
