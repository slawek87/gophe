package settings

type Validation struct {}


func (validation *Validation) isValid(key string, value string) bool {
	const EMPTY = ""

	if key == EMPTY || value == EMPTY {
		return false
	}
	return true
}
