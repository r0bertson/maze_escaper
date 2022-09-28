package maze

import (
	"encoding/json"
	"errors"
)

const (
	InvalidMazeErrorMsg = "this is not a maze"
	NonExistentExitMsg  = "Sorry"
)

// Maze holds the initial position of a maze, which is represented as a JSON-like object key-value object chaining (map[string]interface{}).
type Maze struct {
	InitialRoom Room
}

// FindExitPath looks for the exit. Returns "Sorry" if there isn't one.
func (m *Maze) FindExitPath() interface{} {
	if result := m.InitialRoom.FindExit(); len(result) > 0 {
		return result
	}
	return NonExistentExitMsg
}

// LoadMaze transforms a json-like maze into a named map[string]interface{} called Maze.
// If this map is empty, returns an error.
func LoadMaze(input string) (*Maze, error) {
	var maze Maze
	if err := json.Unmarshal([]byte(input), &maze.InitialRoom); err != nil {
		return nil, err
	}
	if len(maze.InitialRoom) == 0 {
		return nil, errors.New(InvalidMazeErrorMsg)
	}
	return &maze, nil
}
