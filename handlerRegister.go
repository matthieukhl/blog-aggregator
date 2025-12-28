package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	username := cmd.args[0]

	if username == "" {
		return fmt.Errorf(ErrMissingUsername)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Printf("failed to create user %s\n", username)
		os.Exit(1)
	}

	s.cfg.SetUser(username)
	fmt.Printf("User %s created!\n", username)
	fmt.Printf("ID: %v\nName: %s\nCreation Date: %v\nLast Updated: %v\n", user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return nil
}
