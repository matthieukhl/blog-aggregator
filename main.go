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
	ErrFeedDoesNotExist     = "Feed does not exist"
	ErrNoExistingFeeds      = "No feeds registered yet. Register a feed using the 'addFeed' command."
	ErrNoExistingPosts      = "No posts registered yet. Register a post using the 'agg' command."

	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
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
	cmds.register("agg", middlewareLoggedIn(handlerAgg))
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("scrapeFeeds", middlewareLoggedIn(handlerScrapeFeeds))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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
