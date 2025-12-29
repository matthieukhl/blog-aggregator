package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/matthieukhl/blog-aggregator/internal/database"
)

// Handles addfeed command.
// Command takes two arguments:
// - name: the name of the RSS feed
// - url: the URL of the feed
func handlerAddFeed(s *state, cmd command) error {
	// Check if there are enough arguments
	if len(cmd.args) < 2 {
		return fmt.Errorf(ErrNotEnoughArgs)
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf(ErrNoUsersFound)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("Successfully added new feed!")
	fmt.Printf("Feed ID: %v\nFeed name: %s\nFeed URL: %s\nUser ID: %v\nCreated at: %v\nUpdated at: %v\n", result.ID, result.Name, result.Url, result.UserID, result.CreatedAt, result.UpdatedAt)

	return nil
}
