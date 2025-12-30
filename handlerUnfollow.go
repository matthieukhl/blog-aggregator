package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf(ErrNotEnoughArgs)
	}

	feedUrl := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf(ErrFeedDoesNotExist)
		}
		return err
	}

	params := database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}

	return s.db.DeleteFeedFollow(context.Background(), params)

}
