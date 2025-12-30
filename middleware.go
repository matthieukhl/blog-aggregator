package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/matthieukhl/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf(ErrNoUsersFound)
			}
			return err
		}
		return handler(s, cmd, user)
	}
}
