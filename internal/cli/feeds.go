package cli

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.GetFeedsWithUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, f := range feeds {
		fmt.Printf("%s (%s) - added by %s\n", f.FeedName, f.Url, f.UserName)
	}

	return nil
}
