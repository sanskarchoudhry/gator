package cli

import (
	"context"
	"fmt"
)

// HandlerReset clears all users from the database
func HandlerReset(s *State, cmd Command) error {
	err := s.DB.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}
	fmt.Println("✅ Database reset: all users deleted.")
	return nil
}
