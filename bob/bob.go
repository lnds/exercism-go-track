package bob

import (
	"strings"
	"unicode"
)

func asking(str string) bool {
	return strings.HasSuffix(str, "?")
}

func yelling(str string) bool {
	hasLetters := strings.IndexFunc(str, unicode.IsLetter) >= 0
	isUppercase := strings.ToUpper(str) == str
	return hasLetters && isUppercase
}

func saying_anything(str string) bool {
	return str == ""
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var str = strings.TrimSpace(remark)
	switch {
	case asking(str) && yelling(remark):
		return "Calm down, I know what I'm doing!"
	case yelling(str):
		return "Whoa, chill out!"
	case asking(str):
		return "Sure."
	case saying_anything(str):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
