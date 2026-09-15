package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/push-basic/Gator/internal/database"
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

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		return errors.New("No feeds found in database.")
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {

		fmt.Printf("%v, %v, created by: %v\n", feed.Name, feed.Url, feed.UserName)
	}

	return nil

}
