package encode

import (
	"fmt"
	"unicode"
)

func RunLengthEncode(input string) string {
	result := ""
	count := 0
	var char rune
	for _, c := range input {
		if c != char {
			result += writeChar(count, char)
			char = c
			count = 0
		}
		count++
	}
	result += writeChar(count, char)
	return result
}

func writeChar(count int, char rune) string {
	if char == 0 {
		return ""
	}
	if count > 1 {
		return fmt.Sprintf("%d%c", count, char)
	}
	return string(char)
}

func RunLengthDecode(input string) string {
	count := 0
	result := ""
	for _, c := range input {
		if unicode.IsDigit(c) {
			count = count*10 + int(c-'0')
		} else {
			if count > 0 {
				for i := 0; i < count; i++ {
					result += string(c)
				}
			} else {
				result += string(c)
			}
			count = 0
		}
	}
	return result
}
