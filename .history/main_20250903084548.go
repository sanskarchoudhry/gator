package main

import (
	"fmt"
	"gator/internal/config"
	"log"
)

func main() {
	// 1. Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// 2. Set current user to your name
	err = cfg.SetUser("sanskar") // replace with your name
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	// 3. Read again to confirm changes
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading updated config: %v", err)
	}

	// 4. Print the struct
	fmt.Printf("Config: %+v\n", updatedCfg)
}
