package main

import (
	"fmt"

	"github.com/matthieukhl/blog-aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("matthieu")
	cfg = config.Read()
	fmt.Printf("%s\n%s\n", cfg.DBUrl, cfg.CurrentUserName)
}
