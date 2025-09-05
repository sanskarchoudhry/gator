package cli

import (
	"context"
	"fmt"

	"github.com/sanskarchoudhry/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	// Print the parsed feed
	fmt.Println("Channel:", feed.Channel.Title)
	fmt.Println("Description:", feed.Channel.Description)
	fmt.Println("Link:", feed.Channel.Link)
	fmt.Println("\nItems:")

	for _, item := range feed.Channel.Items {
		fmt.Println(" -", item.Title)
	}

	return nil
}
