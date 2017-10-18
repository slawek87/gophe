package settings

type Validation struct {}

// Variables cannot be assigned to empty values.
// Key cannot contains white spaces.
func (validation *Validation) isValid(key string, value string) bool {
	const EMPTY = ""
	const WHITE_SPACE = " "

	if key == EMPTY || value == EMPTY {
		return false
	}

	if key == WHITE_SPACE {
		return false
	}
	return true
}
