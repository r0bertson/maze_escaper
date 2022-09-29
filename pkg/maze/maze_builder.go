package maze

import (
	"github.com/r0bertson/maze_escaper/pkg/utils"
	"github.com/rs/zerolog/log"
	"math/rand"
)

var DefaultDirections = []string{"forward", "right", "left", "upstairs"}
var DefaultObstacles = []string{"dragon", "demon", "lion", "dead end"}

// Builder holds all configurations needed to create a random maze.
type Builder struct {
	MaxDepth         int
	ObstacleRate     float64
	PathWideningRate float64
	Directions       utils.Tokens
	Obstacles        utils.Tokens
	ExitsRemaining   int
}

func NewBuilder(maxDepth *int, obstacleRate, pathWideningRate *float64, obstaclesFilepath, directionsFilepath *string) Builder {
	return Builder{
		ExitsRemaining:   1,
		MaxDepth:         *maxDepth,
		ObstacleRate:     *obstacleRate,
		PathWideningRate: *pathWideningRate,
		Directions:       LoadInput(directionsFilepath, DefaultDirections),
		Obstacles:        LoadInput(obstaclesFilepath, DefaultObstacles),
	}
}

func (mb *Builder) GetDirectionRandomDestination(depth int) interface{} {
	if rand.Float64() < 0.01 && mb.ExitsRemaining > 0 {
		mb.ExitsRemaining -= 1
		return "exit"
	}
	if rand.Float64() < mb.PathWideningRate && depth <= mb.MaxDepth {
		return mb.GenerateRandomPath(depth + 1)
	}
	if rand.Float64() < mb.ObstacleRate {
		return mb.Obstacles.GetRandom()
	}
	return nil
}

// GenerateRandomMaze generates a randomized maze.
// Depending on MaxDepth, there is no guarantee that the maze will have an exit
func (mb *Builder) GenerateRandomMaze() *Maze {
	var initialPosition Path = map[string]interface{}{}
	for _, direction := range mb.Directions {
		initialPosition[direction] = mb.GenerateRandomPath(1)
	}
	return &Maze{Start: initialPosition}
}

func (mb *Builder) GenerateRandomPath(depth int) Path {
	var room Path = map[string]interface{}{}
	for _, direction := range mb.Directions {
		destination := mb.GetDirectionRandomDestination(depth)
		switch destination.(type) {
		case nil:
			//skip direction
			continue
		case string:
			//add obstacle
			room[direction] = destination
		default:
			//generate a path as child
			room[direction] = destination
		}
	}

	return room
}

func LoadInput(filepath *string, defaultData []string) utils.Tokens {
	if filepath == nil {
		return defaultData
	}
	loaded, err := utils.ReadTokensFromFile(*filepath)
	if err != nil {
		log.Warn().Msgf("couldn't read input file (%s) (error: %v) using default instead", *filepath, err)
		return defaultData
	}
	return loaded
}
