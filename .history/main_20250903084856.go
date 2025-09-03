package main

import (
	"fmt"
	"log"

	"github.com/sanskarchoudhry/gator/internal/config"
)

func main() {
	// 1. Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// 2. Set current user to your name
	err = cfg.SetUser("jojo")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	// 3. Read again to confirm changes
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading updated config: %v", err)
	}

	fmt.Printf("Config: %+v\n", updatedCfg)
}
