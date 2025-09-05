package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("a username is required")
	}
	username := cmd.Args[0]

	// Ensure user exists
	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no such user %q", username)
		}
		return fmt.Errorf("looking up user: %w", err)
	}

	if err := s.Cfg.SetUser(username); err != nil {
		return fmt.Errorf("saving current user: %w", err)
	}

	fmt.Printf("logged in as %s\n", username)
	return nil
}
