package main

import (
	"context"
	"fmt"
	"landovargas/blog-aggregator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	var limit int
	var err error
	if len(cmd.Args) == 0 {
		limit = 2
	} else {

		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Failed to conver: %v to number, err: %w", cmd.Args[0], err)
		}
	}

	posts, err := s.db.GetPostForUser(context.Background(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("Failed to retrieve posts from database")
	}

	for _, post := range posts {
		printPost(post)
	}

	return nil
}

func printPost(post database.Post) {
	fmt.Println("====================")
	fmt.Printf("* PubDate:       %s\n", post.PublishedAt.Time)
	fmt.Printf("* Title:         %s\n", post.Title)
	fmt.Printf("* Content:       %s\n", post.Description.String)
	fmt.Printf("* URL:           %s\n", post.Url)

}
