package main

import "fmt"

type commands struct {
	functions map[string]func(*state, command) error
}

func NewCommands() commands {
	return commands{
		functions: make(map[string]func(*state, command) error),
	}
}

// This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	commandToRun, ok := c.functions[cmd.name]
	if !ok {
		return fmt.Errorf(ErrCommandDoesNotExists)
	}

	return commandToRun(s, cmd)
}

// This command registters a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	c.functions[name] = f
}
