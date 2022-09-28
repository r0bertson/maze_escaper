package maze

import "github.com/rs/zerolog/log"

// Room is a named type for a map[string]interface{} to allow adding extended functionality to a composite type.
type Room map[string]interface{}

/*
FindExit tires to find an exit's path from a certain room. It assumes that each maze's nodes (rooms) can only be
string or map[string]interface{}.
This is a recursion, so it'll spawn all possibilities moving down a json-like structure (maze's expected input
format) until it reaches an exit or look into every possibility. This is not obeying any laws of physics, so going
back to a previous room is not possible (e.g. going "left" 3 times won't return to the same room).
If this feature is desired, we should move away from this json-like structure.
*/
func (r *Room) FindExit() []string {
	for direction, value := range *r {
		switch t := value.(type) {
		case string:
			if value == "exit" {
				return []string{direction}
			}
			continue //found obstacle, abandon this route
		case map[string]interface{}:
			var nextRoom Room = value.(map[string]interface{})
			if nextRoom != nil {
				path := nextRoom.FindExit()
				if len(path) > 0 {
					//if result is not empty, an exit was found
					return append([]string{direction}, path...)
				}
			}
		default:
			log.Info().Msgf("unexpected room representation %v", t)
			continue
		}

	}
	//no exits found
	return []string{}
}
