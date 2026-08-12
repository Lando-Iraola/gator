package main

import (
	"context"
	"fmt"
	"landovargas/blog-aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to find user: %v", err)
	}

	var feed database.CreateFeedParams

	feed.ID = uuid.New()
	feed.CreatedAt = time.Now()
	feed.UpdatedAt = time.Now()
	feed.Name = name
	feed.Url = url
	feed.UserID = user.ID
	newFeed, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return fmt.Errorf("Failed to create new feed: %w", err)
	}

	fmt.Printf("%v", newFeed)

	return nil
}
