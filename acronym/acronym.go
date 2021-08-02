package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) (out string) {
	var parts = regexp.MustCompile("[a-zA-Z']+").FindAllString(s, -1)
	for i := range parts {
		out += string(parts[i][0])
	}
	return strings.ToUpper(out)
}
