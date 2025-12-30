package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func handlerScrapeFeeds(s *state, cmd command, user database.User) error {
	// Get next feed to fetch from db
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf(ErrNoExistingFeeds)
		}
		return err
	}

	// Mark feed as fetched
	params := database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:     time.Now(),
	}
	err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf(ErrFeedDoesNotExist)
		}
		return err
	}

	// Query feed URL to get feed data
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	// Print item titles to console
	for i, item := range feedData.Channel.Item {
		fmt.Printf("%d: %s\n", i+1, item.Title)
	}

	return nil
}
