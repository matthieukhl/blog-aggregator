package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		fmt.Printf("failed to create new GET reequest to %s:%v\n", feedURL, err)
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to execute request to %s:%v\n", feedURL, err)
		return nil, err
	}

	// Unmarshal response
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	// Decode escaped HTML entities
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for index, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[index].Title = html.UnescapeString(item.Title)
		rssFeed.Channel.Item[index].Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil

}
