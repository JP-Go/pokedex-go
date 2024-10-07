package main

import (
	"time"

	"github.com/JP-Go/pokedex-go/internal/cache"
	"github.com/JP-Go/pokedex-go/internal/commands"
	"github.com/JP-Go/pokedex-go/internal/repl"
)

func main() {
	cacheConfig := cache.NewCache(5 * time.Minute)
	cliConfig := commands.NewCliConfig(cacheConfig)
	repl.StartRepl(cliConfig)
}
