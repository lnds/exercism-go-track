package atbash

import "strings"

func Atbash(plain string) string {
	cipher := ""
	group := 0
	for _, c := range strings.ToLower(plain) {
		if c >= 'a' && c <= 'z' {
			cipher += string(rune('z' - (c - 'a')))
			group++
		} else if c >= '0' && c <= '9' {
			cipher += string(c)
			group++
		}
		if group == 5 {
			cipher += " "
			group = 0
		}
	}
	return strings.TrimSpace(cipher)
}
