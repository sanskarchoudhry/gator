package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sanskarchoudhry/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	// Get current user from config
	currentUser, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not find current user: %w", err)
	}

	// Create new feed
	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("create feed: %w", err)
	}

	// Print result
	log.Printf("Feed created: %s (%s) by user %s", feed.Name, feed.Url, currentUser.Name)
	fmt.Println("Feed ID:", feed.ID)
	fmt.Println("Created At:", feed.CreatedAt)
	fmt.Println("Updated At:", feed.UpdatedAt)
	fmt.Println("Name:", feed.Name)
	fmt.Println("URL:", feed.Url)
	fmt.Println("UserID:", feed.UserID)

	return nil
}
