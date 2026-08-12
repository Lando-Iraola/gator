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

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	req.Header.Add("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to get feed, status code: %v, status: %v", resp.StatusCode, resp.Status)
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to read feed: %w", err)
	}

	var rss RSSFeed
	err = xml.Unmarshal(b, &rss)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse feed: %w", err)
	}

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Title)

	return &rss, nil
}

func handlerRss(s *state, cmd command) error {
	/*if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <link>", cmd.Name)
	}

	link := cmd.Args[0]*/

	link := "https://www.wagslane.dev/index.xml"

	duration, err := time.ParseDuration("30s")
	if err != nil {
		return fmt.Errorf("Failed to create rss context duration: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	rss, err := fetchFeed(ctx, link)

	if err != nil {
		return err
	}

	fmt.Printf("%v", rss)

	return nil
}
