package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sanskarchoudhry/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: follow <feed-url>")
	}

	feedURL := cmd.Args[0]

	// Get current user
	user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}

	// Find feed by URL
	feed, err := s.DB.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("feed not found: %w", err)
	}

	// Create follow record
	now := time.Now()
	follow, err := s.DB.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			UserID:    user.ID,
			FeedID:    feed.ID,
		})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("User %s is now following feed %s\n", follow.UserName, follow.FeedName)
	return nil
}
