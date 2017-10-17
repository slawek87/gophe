package settings

type Validation struct {}

// checks if current key and value are valid settings - cannot be empty strings.
func (validation *Validation) isValid(key string, value string) bool {
	if key == "" || value == "" {
		return false
	}
	return true
}
