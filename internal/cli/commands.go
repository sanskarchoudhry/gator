package cli

import (
	"fmt"

	"github.com/sanskarchoudhry/gator/internal/config"
	"github.com/sanskarchoudhry/gator/internal/database"
)

type Command struct {
	Name string
	Args []string
}

type State struct {
	Cfg *config.Config
	DB  *database.Queries
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
