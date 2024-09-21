package main

import (
	"github.com/JP-Go/pokedex-go/internal/commands"
	"github.com/JP-Go/pokedex-go/internal/repl"
)

func main() {
	repl.StartRepl(commands.NewCliConfig())
}
