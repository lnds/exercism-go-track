package pangram

import "strings"

func IsPangram(word string) bool {
	word = strings.ToLower(word)
	chars := map[rune]bool{}
	for _, r := range word {
		if r >= 'a' && r <= 'z' {
			chars[r] = true
		}
	}
	return len(chars) == 26
}
