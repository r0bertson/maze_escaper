package main

import (
	"encoding/json"
	"flag"
	"github.com/r0bertson/maze_escaper/pkg/maze"
	"github.com/r0bertson/maze_escaper/pkg/utils"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	maxDepth := flag.Int("d", 10, "max depth of maze tree")
	obstacleRate := flag.Float64("or", 0.4, "obstacle occurrence rate")
	pathWideningRate := flag.Float64("pr", 0.2, "path widening rate")
	obsFilepath := flag.String("obs", "../../pkg/seed/obstacles.txt", "path to obstacles file")
	dirFilepath := flag.String("dir", "../../pkg/seed/directions.txt", "path to directions file")
	specialDirFilepath := flag.String("spe", "../../pkg/seed/special_directions.txt", "path to special directions file")
	export := flag.Bool("export", false, "export executions to file")
	flag.Parse()
	builder := maze.NewBuilder(maxDepth, obstacleRate, pathWideningRate, obsFilepath, dirFilepath, specialDirFilepath)

	random := builder.GenerateRandomMaze()
	solution := random.FindExitPath()

	now := utils.GetFileTimestamp()
	if *export {
		mazeJson, err := json.MarshalIndent(random.Start, "", "    ")
		if err != nil {
			log.Warn().Err(err)
		}
		resultJson, err := json.MarshalIndent(solution, "", "    ")
		if err != nil {
			log.Warn().Err(err)
		}
		if err = os.WriteFile(string(now)+"_maze.json", mazeJson, 0644); err != nil {
			log.Warn().Err(err)
		}
		if err = os.WriteFile(string(now)+"_solution.json", resultJson, 0644); err != nil {
			log.Warn().Err(err)
		}
	}

}
