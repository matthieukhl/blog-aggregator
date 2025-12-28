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