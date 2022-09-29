package maze

import (
	"github.com/r0bertson/maze_escaper/pkg/utils"
	"testing"
)

var validEmptyPath Path = map[string]interface{}{}
var validPathWithObstacle Path = map[string]interface{}{"forward": "lion"}
var validPathWithPathsButNoExit Path = map[string]interface{}{"forward": map[string]interface{}{"left": "lion"}, "right": "dragon"}
var validPathWithPathsAndExit Path = map[string]interface{}{"forward": map[string]interface{}{"left": "lion"}, "right": "exit"}
var validPathWithWithInvalidStructure Path = map[string]interface{}{"forward": 123}
var emptyPath []string

// TestPath_FindExit will try to find a path to an exit starting from a certain location.
func TestPath_FindExit(t *testing.T) {
	samples := []struct {
		Path     Path
		Expected []string
	}{
		{validEmptyPath, emptyPath},
		{validPathWithObstacle, emptyPath},
		{validPathWithPathsButNoExit, emptyPath},
		{validPathWithPathsAndExit, []string{"right"}},
		{validPathWithWithInvalidStructure, emptyPath},
	}

	for idx, input := range samples {
		if path := input.Path.FindExit(); !utils.IsEqual(path, input.Expected) {
			t.Errorf("FindExit: sample idx %d has a result %v different than expected %v", idx, path, input.Expected)

		}
	}
}
