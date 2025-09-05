package cli

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	// Get current user
	user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	for _, f := range follows {
		fmt.Printf("* %s\n", f.FeedName)
	}

	return nil
}
