package maze

// Path is a named type for a map[string]interface{} to allow adding extended functionality to a composite type.
type Path map[string]interface{}

/*
FindExit tires to find an exit. It assumes that each maze's nodes (paths) can only be string or map[string]interface{}.
This is a recursion, so it'll spawn all possibilities moving down a json-like structure (maze's expected input
format) until it reaches an exit or look into every possibility. This is not obeying any laws of physics, so going
back to a previous location is not possible (e.g. going "left" 4 times won't return to the initial position).
If this feature is desired, we should move away from this json-like structure.
*/
func (r *Path) FindExit() []string {
	for direction, value := range *r {
		switch value.(type) {
		case string:
			if value == "exit" {
				return []string{direction}
			}
			continue //found obstacle, abandon this route
		case map[string]interface{}:
			//sometimes casting map[string]interface{} to Path throws and error
			//this is handled by another case, but needs investigation
			var nextPath Path = value.(map[string]interface{})
			if nextPath != nil {
				path := nextPath.FindExit()
				if len(path) > 0 {
					//if result is not empty, an exit was found
					return append([]string{direction}, path...)
				}
			}
		case Path:
			if nextPath := value.(Path); nextPath != nil {
				path := nextPath.FindExit()
				if len(path) > 0 {
					//if result is not empty, an exit was found
					return append([]string{direction}, path...)
				}
			}
		default:
			//log.Info().Msgf("unexpected path representation %v", t)
			continue
		}

	}
	//no exits found
	return []string{}
}
