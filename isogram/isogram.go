package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(text string) bool {
	var letters = map[rune]bool{}
	for _, c := range strings.ToLower(text) {
		if unicode.IsLetter(c) {
			if letters[c] {
				return false
			}
			letters[c] = true
		}
	}
	return true
}
