package settings

import "strings"

type Comments struct {}

// checks if given string is a comment - starts with: /, //, #.
func (comments *Comments) isComment(text string) bool {
	commentsPattern := []string{"/", "//", "#"}

	for _, element := range commentsPattern {
		if strings.HasPrefix(text, element) == true {
			return true
		}
	}

	return false
}
