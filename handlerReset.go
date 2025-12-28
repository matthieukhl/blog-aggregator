package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	_, err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println(ErrFailedToDeleteUsers)
		os.Exit(1)
	}
	return nil
}
