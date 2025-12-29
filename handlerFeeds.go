package main

import (
	"context"
	"fmt"
)

// Handles feed command
// Fetches all feeds from feeds table and prints them to stdout
func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println("===================")
		fmt.Printf("Feed name: %s\nFeed URL: %s\nCreated by: %s\n", feed.Name, feed.Url, feed.Username)
		fmt.Println("===================")
	}

	return nil
}
