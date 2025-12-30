package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf(ErrNotEnoughArgs)
	}

	feedUrl := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf(ErrFeedDoesNotExist)
	}

	// user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return fmt.Errorf(ErrNoUsersFound)
	// 	}
	// 	return err
	// }

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println("==================")
	fmt.Printf("User %s is now subscribed to %q feed!\n", result.UserName, result.FeedName)
	fmt.Println("==================")

	return nil
}
