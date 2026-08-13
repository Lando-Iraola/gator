package main

import (
	"context"
	"fmt"
	"landovargas/blog-aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFeedFollow(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.FeedByURL(context.Background(), url)

	if err != nil {
		return fmt.Errorf("Failed to find feed by url: %v, err: %w", url, err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Failed to create feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(follow.User, follow.Feed)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("Failed to find follows for user: %v, err: %w", user.Name, err)
	}

	if len(follows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range follows {
		fmt.Printf("* %s\n", ff.Feed)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
