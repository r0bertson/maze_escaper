package maze

import (
	"github.com/r0bertson/maze_escaper/pkg/utils"
	"testing"
)

var validEmptyRoom Room = map[string]interface{}{}
var validRoomWithObstacle Room = map[string]interface{}{"forward": "lion"}
var validRoomWithPathsButNoExit Room = map[string]interface{}{"forward": map[string]interface{}{"left": "lion"}, "right": "dragon"}
var validRoomWithPathsAndExit Room = map[string]interface{}{"forward": map[string]interface{}{"left": "lion"}, "right": "exit"}
var validRoomWithWithInvalidStructure Room = map[string]interface{}{"forward": 123}
var emptyPath []string

// TestRoom_FindExit will try to find a path to an exit starting from a room.
func TestRoom_FindExit(t *testing.T) {
	samples := []struct {
		Room     Room
		Expected []string
	}{
		{validEmptyRoom, emptyPath},
		{validRoomWithObstacle, emptyPath},
		{validRoomWithPathsButNoExit, emptyPath},
		{validRoomWithPathsAndExit, []string{"right"}},
		{validRoomWithWithInvalidStructure, emptyPath},
	}

	for idx, input := range samples {
		if path := input.Room.FindExit(); !utils.IsEqual(path, input.Expected) {
			t.Errorf("FindExit: sample idx %d has a result %v different than expected %v", idx, path, input.Expected)

		}
	}
}
