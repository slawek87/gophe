package settings

type Validation struct {}

// checks if current key and value are valid settings - cannot be empty strings.
func (validation *Validation) isValid(key string, value string) bool {
	const EMPTY = ""

	if key == EMPTY || value == EMPTY {
		return false
	}
	return true
}
