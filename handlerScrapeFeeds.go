package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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

	// Insert posts in posts table
	newPostsAddedCount := 0
	for _, item := range feedData.Channel.Item {
		// TODO: refactor code using a helper function to create postParams.
		params := database.CreatePostParams{}
		parsedPubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			params = database.CreatePostParams{
				ID:          uuid.New(),
				Title:       item.Title,
				Description: sql.NullString{String: item.Description, Valid: true},
				Url:         item.Link,
				PublishedAt: sql.NullTime{Time: time.Time{}, Valid: false},
				FeedID:      feed.ID,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
		} else {
			params = database.CreatePostParams{
				ID:          uuid.New(),
				Title:       item.Title,
				Description: sql.NullString{String: item.Description, Valid: true},
				Url:         item.Link,
				PublishedAt: sql.NullTime{Time: parsedPubDate, Valid: true},
				FeedID:      feed.ID,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
		}

		result, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			return err
		}
		newPostsAddedCount += 1
		fmt.Printf("New post: %s!\n", result.Title)
	}

	fmt.Printf("\nSuccessfully inserted %d new posts from %s!\n", newPostsAddedCount, feed.Name)

	return nil
}
