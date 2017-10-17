package settings

type Validation struct {}

// Variables cannot be assigned to empty values.
func (validation *Validation) isValid(key string, value string) bool {
	const EMPTY = ""

	if key == EMPTY || value == EMPTY {
		return false
	}
	return true
}
