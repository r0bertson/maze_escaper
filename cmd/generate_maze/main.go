package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/r0bertson/maze_escaper/pkg/maze"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	builder := maze.NewBuilder(
		flag.Int("d", 5, "max depth of maze tree"),
		flag.Float64("pr", 0.4, "obstacle occurrence rate"),
		flag.Float64("or", 0.2, "path widening rate"),
		flag.String("obs", "./pkg/seed/obstacles.txt", "path to obstacles file"),
		flag.String("dir", "./pkg/seed/directions.txt", "path to directions file"),
		flag.String("spe", "./pkg/seed/special_directions.txt", "path to special directions file"),
	)
	random := builder.GenerateRandomMaze()
	solution := random.FindExitPath()
	fmt.Println(solution)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	if err := enc.Encode(random.Start); err != nil {
		log.Fatal().Err(err)
	}

}
