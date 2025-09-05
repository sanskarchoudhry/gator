package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sanskarchoudhry/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("a username is required")
	}
	name := cmd.Args[0]

	// If user already exists, exit with error
	_, err := s.DB.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user %q already exists", name)
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("checking existing user: %w", err)
	}

	now := time.Now().UTC()

	// NOTE: sqlc typically generates a Params struct; adjust if your signature differs.
	u, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	// Set current user in config
	if err := s.Cfg.SetUser(name); err != nil {
		return fmt.Errorf("saving current user: %w", err)
	}

	fmt.Printf("✅ user %q created and set as current\n", name)
	fmt.Printf("user: id=%s created_at=%s updated_at=%s name=%s\n",
		u.ID, u.CreatedAt.Format(time.RFC3339), u.UpdatedAt.Format(time.RFC3339), u.Name)
	return nil
}
