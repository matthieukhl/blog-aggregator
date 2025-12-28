package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf(ErrNoArgs)
	}

	username := cmd.args[0]
	exists, err := isUserExists(s, username)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Println(ErrUserDoesNotExists)
		os.Exit(1)
	}

	s.cfg.SetUser(username)

	fmt.Printf("User %s has been set succesfully.\n", username)

	return nil
}

func isUserExists(s *state, username string) (bool, error) {
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf(ErrUserDoesNotExists)
		}
		return false, err
	}

	return true, nil
}
