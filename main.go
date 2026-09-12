package main

import (
	"Gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("push")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
}
