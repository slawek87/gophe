package settings

import "strings"

type Comments struct {}

// Use that method to check if given string is a comment.
func (comments *Comments) isComment(configLine string) bool {
	commentsPattern := []string{"/", "//", "#"}

	for _, comment := range commentsPattern {
		if strings.HasPrefix(configLine, comment) == true {
			return true
		}
	}

	return false
}
