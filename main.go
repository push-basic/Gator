package main

import (
	"Gator/internal/config"
	"fmt"
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

	fmt.Println("URL:", cfg.URL)
	fmt.Println("Name:", cfg.Name)
}
