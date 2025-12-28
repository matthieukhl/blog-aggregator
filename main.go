package main

import (
	"fmt"
	"os"

	"github.com/matthieukhl/blog-aggregator/internal/config"
)

const (
	ErrNoArgs               = "No arguments set to command"
	ErrCommandDoesNotExists = "Command does not exists"
	ErrNotEnoughArgs        = "Not enough arguments."
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func NewCommand(name string, args []string) command {
	return command{
		name: name,
		args: args,
	}
}

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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf(ErrNoArgs)
	}

	fmt.Println(len(cmd.args))

	username := cmd.args[0]
	s.cfg.SetUser(username)

	fmt.Printf("User %s has been set succesfully.\n", username)

	return nil
}

func main() {
	cfg := config.Read()
	newState := state{
		cfg: cfg,
	}
	cmds := NewCommands()
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println(ErrNotEnoughArgs)
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := NewCommand(cmdName, cmdArgs)

	err := cmds.run(&newState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
