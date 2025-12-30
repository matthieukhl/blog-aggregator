package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	feedsFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User is not subscribed to a feed yet! Subscribe to your first feed with the follow command")
		}
		return err
	}

	fmt.Println("==================")
	fmt.Println("Currently subscribed to:")
	for index, feed := range feedsFollowing {
		fmt.Printf("\t%d: %s\n", index+1, feed.FeedName)
	}
	fmt.Println("==================")

	return nil
}
