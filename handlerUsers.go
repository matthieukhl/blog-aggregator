package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(ErrNoUsersFound)
			os.Exit(1)
		}
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}

	}

	return nil
}
