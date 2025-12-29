package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/matthieukhl/blog-aggregator/internal/config"
	"github.com/matthieukhl/blog-aggregator/internal/database"
)

const (
	ErrNoArgs               = "No arguments set to command"
	ErrCommandDoesNotExists = "Command does not exists"
	ErrNotEnoughArgs        = "Not enough arguments."
	ErrMissingUsername      = "Missing username"
	ErrUserDoesNotExists    = "User does not exists"
	ErrFailedToDeleteUsers  = "Failed to delete users"
	ErrNoUsersFound         = "No users found"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg := config.Read()
	newState := state{
		cfg: cfg,
	}
	db, err := sql.Open("postgres", newState.cfg.DBUrl)
	dbQueries := database.New(db)
	newState.db = dbQueries
	cmds := NewCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

	if len(os.Args) < 2 {
		fmt.Println(ErrNotEnoughArgs)
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := NewCommand(cmdName, cmdArgs)

	err = cmds.run(&newState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
