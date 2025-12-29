package main

import "fmt"

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf(ErrNotEnoughArgs)
	}

	feedUrl := cmd.args[0]

}
