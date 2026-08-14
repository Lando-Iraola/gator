package main

import (
	"context"
	"database/sql"
	"fmt"
	"landovargas/blog-aggregator/internal/database"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenReqsText := cmd.Args[0]
	timeBetweenReqs, err := time.ParseDuration(timeBetweenReqsText)
	if err != nil {
		return fmt.Errorf("Failed to set time between requests: %w", err)
	}

	fmt.Println(fmt.Sprintf("Collecting feeds every %s", timeBetweenReqs.String()))
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		var desc sql.NullString
		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}

		tim, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			tim, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				log.Printf("Failed to format post time: %v", err)
			}
		}
		var pubDate sql.NullTime
		pubDate.Valid = false
		if err == nil {
			pubDate.Time = tim
			pubDate.Valid = true
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "23505") {
				continue
			}
			log.Printf("Error encountered: %w", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
