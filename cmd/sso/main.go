package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger

	// TODO: init app

	// TODO: run app
}
