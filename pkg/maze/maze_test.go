package maze

import (
	"github.com/r0bertson/maze_escaper/pkg/utils"
	"testing"
)

type Sample struct {
	Input    string
	Expected interface{}
}

var findExitSamples = []Sample{
	{
		Input:    firstInputFromPDF,
		Expected: []string{"left", "forward", "upstairs"},
	},
	{
		Input:    secondInputFromPDF,
		Expected: []string{"forward"},
	},
	{
		Input:    thirdInputFromPDF,
		Expected: "Sorry",
	},
	{
		Input:    mazeWithInvalidStructure,
		Expected: "Sorry",
	},
}

const (
	firstInputFromPDF  = `{ "forward": "tiger", "left": {"forward": { "upstairs": "exit" }, "left": "dragon"}, "right": {"forward": "dead end"} }`
	secondInputFromPDF = `{ "forward": "exit" }`
	thirdInputFromPDF  = `{ "forward": "tiger", "left": "ogre", "right": "demon" }`

	emptyMaze                = `{}`
	mazeWithInvalidStructure = `{ "forward": [1,2,3], "left": 3.1415, "right": 134 }`
)

var loadMazeSamplesFromPDF = []string{firstInputFromPDF, secondInputFromPDF, thirdInputFromPDF}

// TestLoadMaze_PDFSamples testes if all samples are loaded and valid (initial room has at least one `door` to move)
func TestLoadMaze_PDFSamples(t *testing.T) {
	for idx, input := range loadMazeSamplesFromPDF {
		loadedMaze, err := LoadMaze(input)
		if err != nil || len(loadedMaze.InitialRoom) == 0 {
			t.Errorf("unable to load sample maze index at %d, error was: %v", idx, err)
		}
	}
}

// TestLoadMaze_EmptyInput will try to load an empty representation of a Maze `{}`.
// Since this is an initial room without doors to move, it is considered invalid and an error is expected.
func TestLoadMaze_EmptyInput(t *testing.T) {
	_, err := LoadMaze(emptyMaze)
	if err == nil {
		t.Errorf("maze was loaded, but should thrown an error")
	}
}

// TestLoadMaze_InvalidStructure will try to load a maze with invalid structure (room represented by invalid types)
// it should load without any issue and invalid rooms will be ignored when trying to find an exit.
func TestLoadMaze_InvalidStructure(t *testing.T) {
	loadedMaze, err := LoadMaze(mazeWithInvalidStructure)
	if err != nil || len(loadedMaze.InitialRoom) == 0 {
		t.Errorf("maze wasn't loaded properly. invalid structure shouldn't affect maze loading")
	}
}

// TestMaze_FindExitPath checks if the path found by the FindExitPath function is the one expected.
func TestMaze_FindExitPath(t *testing.T) {
	for idx, input := range findExitSamples {
		loadedMaze, _ := LoadMaze(input.Input)
		result := loadedMaze.FindExitPath()
		switch result.(type) {
		case string:
			if result != input.Expected {
				t.Errorf("FindExitPath: sample idx %d has a result %v different than expected: %v", idx, result, input.Expected)
			}
		case []string:
			resultAsSlice := result.([]string)
			expectedAsSlice := input.Expected.([]string)
			if !utils.IsEqual(resultAsSlice, expectedAsSlice) {
				t.Errorf("FindExitPath: sample idx %d has a result %v different than expected %v", idx, resultAsSlice, expectedAsSlice)
			}
		}
	}
}
