package main

import (
	"fmt"
	"time"

	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	// Check if enough args
	if len(cmd.args) < 1 {
		fmt.Println("Usage: agg <time_between_reqs>. time_between_reqs should be a string like '1s', '1m' or '1h'.")
		return fmt.Errorf(ErrNotEnoughArgs)
	}

	timeBetweenReqs := cmd.args[0]
	parsedTimeBetweenReqs, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return fmt.Errorf("Failed to parse argument time_between_requests given %s as input: %v\n", timeBetweenReqs, err)
	}

	fmt.Printf("Collecting feeds every %s\n", parsedTimeBetweenReqs)

	ticker := time.NewTicker(parsedTimeBetweenReqs)
	for range ticker.C {
		handlerScrapeFeeds(s, cmd, user)
	}

	return nil
}
