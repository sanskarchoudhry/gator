package main

import (
	"log"
	"os"

	"github.com/sanskarchoudhry/gator/internal/cli"
	"github.com/sanskarchoudhry/gator/internal/config"
)

func main() {
	// Load config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &cli.State{Cfg: &cfg}

	// Setup command registry
	cmds := cli.Commands{Handlers: make(map[string]func(*cli.State, cli.Command) error)}
	cmds.Register("login", cli.HandlerLogin)

	// Parse command-line args
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:], // everything after "login"
	}

	if err := cmds.Run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
