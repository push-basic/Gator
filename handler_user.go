package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"uuid"

	"github.com/push-basic/Gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, cmd.Args[0])
	if err == nil {
		fmt.Println("user already exists")
		os.Exit(1)
	}

	u, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		fmt.Printf("Failed to create user %v\n", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(u.Name)
	if err != nil {
		return fmt.Errorf("Failed to set current user: %w", err)
	}

	fmt.Printf("User created %v\n", u)

	return nil

}
