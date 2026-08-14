package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time between requests>", cmd.Name)
	}

	timeBetweenReqsText := cmd.Args[0]
	timeBetweenReqs, err := time.ParseDuration(timeBetweenReqsText)
	if err != nil {
		return fmt.Errorf("Failed to set time between requests: %w", err)
	}

	fmt.Printf("Collecting feeds every %s", timeBetweenReqs.String())
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		fmt.Println("Attempting...")
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("Failed to scrape: %w", err)
		}
		fmt.Println("sukesful!")
	}

	return nil
}
