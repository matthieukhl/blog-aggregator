# Blog Aggregator
## Overview
This project is part of boot.dev's guided projects. It's a CLI tool that allows users to:
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Requirements
Make sure you have the following installed on your machine:
1. Go
2. goose
3. sqlc
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
4. PostgreSQL
```bash
# macOS
brew install postgresql@15

# Linux / WSL (Debian)
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo passwd postgres # (Linux only) Update postgres password
```

## Prerequisites
### Setup PostgreSQL database
1. Create gator database
```sql
CREATE DATABASE gator;
```

2. (Linux only) Set the user password
```sql
\c gator
ALTER USER postgres PASSWORD <password>
```

3. Database migration
```bash
goose postgres <connection_string> up

# example:
# goose postgres "postgres://wagslane:@localhost:5432/gator" up
```

### Setup config file
1. Create `.gatorconfig.json` file in your user's home directory.
```bash
touch ~/.gatorconfig.json
```

2. Add the following content to `~/.gatorconfig.json`
```json
{
  "db_url": <your_postgresql_connection_string>
}
```

## Installation
```bash
go install -o gator github.com/matthieukhl/blog-aggregator
```

## Usage
Protected commands can only be executed by logged in users.
- `register <username>`: registers a new user.
- `login <username>`: logs in with specified username. Note: username must be registered before logging in.
- `reset`: deletes all users from users table. Subsequently deletes every row in other tables due to ON DELETE CASCADE constaints.
- `users`: returns a list of all users.
- `agg <interval_between_requests`: (protected) fetches posts from registered feeds. It takes an interval between requests as argument such as 1s, 1m or 1h.
- `addfeed <feed_name> <feed_url>`: (protected) adds a feed to the feeds table and subscribes to it. It takes two arguments: the feed's name and its url.
- `feeds`: returns a list of all subscribed feeds.
- `follow <feed_url>`: (protected) subscribes user to a feed. It takes one argument: the feed's url. Note: feed must first be added to the feeds table.
- `following`: (protected) returns a list of all the feeds the user is subscribed to.
- `unfollow <feed_url>`: (protected) unsubscibes user from specified feed. It takes one argument: the feed's url.
- `browse <limit>`: (protected) returns post title and their description. It takes one optional argument to set the limit of posts returned. If none, defaults to 2.