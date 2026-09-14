package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/push-basic/Gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
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
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Feed created!\n")
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Listing feeds!\n")

	for _, feed := range feeds {
		printFeed(feed)
	}

	return nil

}

func printFeed(feed database.Feed) {
	fmt.Println()
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
	fmt.Println()
}
