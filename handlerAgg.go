package main

import (
	"context"
	"fmt"
	"os"
)

const (
	feedUrl = "https://www.wagslane.dev/index.xml"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(feed)

	return nil
}
