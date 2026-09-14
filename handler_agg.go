package main

import (
	"context"
	"fmt"

	"github.com/push-basic/Gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://hnrss.org/newest")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func addFeedHandler(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Feed created!\n")
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)

	return nil
}
