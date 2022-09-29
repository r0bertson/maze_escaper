package main

import (
	"fmt"
	"github.com/r0bertson/maze_escaper/pkg/maze"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal().Msg("expected maze in json format")
		return
	}
	loaded, err := maze.LoadMaze(args[0])
	if err != nil {
		log.Fatal().Msg("unable to load maze")
	}

	path := loaded.FindExitPath()
	fmt.Println(path)
}
