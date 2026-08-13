package main

import (
	"context"
	"fmt"
	"landovargas/blog-aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.FeedList(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to retrieve feeds: %w", err)
	}
	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, val := range feeds {
		fmt.Println("=====================================")
		printFeeds(val)
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func printFeeds(feeds database.FeedListRow) {

	fmt.Printf("* ID:            %s\n", feeds.ID)
	fmt.Printf("* URL:           %s\n", feeds.Url)
	fmt.Printf("* Feed:          %v\n", feeds.Feed)
	fmt.Printf("* URL:           %v\n", feeds.Url)
	fmt.Printf("* User:          %s\n", feeds.User)
	fmt.Printf("* Created:       %v\n", feeds.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feeds.UpdatedAt)

}
