package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.args) > 0 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf(ErrNoExistingPosts)
		}
		return err
	}

	for i, post := range posts {
		fmt.Printf("\n%d. %s%s%s\n\n\t%s\n\n\t%s\n\n\t%sFetched from: %s\n\tPublication date: %s%sS\n", i+1, Magenta, post.Title, Reset, post.Description.String, post.Url, Yellow, post.FeedName, post.PublishedAt.Time.Format("2006-01-02"), Reset)
	}

	return nil
}
