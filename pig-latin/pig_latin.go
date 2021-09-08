package piglatin

import (
	"strings"
)

func Sentence(input string) string {
	words := strings.Split(input, " ")
	elems := []string{}
	for _, word := range words {
		elems = append(elems, Word(word))
	}
	return strings.Join(elems, " ")
}

func Word(input string) string {
	input, processed := rule1(input)
	if !processed {
		input, processed = rule3(input)
	}
	if !processed {
		input, processed = rule2(input)
	}

	if !processed {
		input, _ = rule4(input)
	}
	return input
}

func rule1(input string) (string, bool) {
	prefixes := []string{"a", "e", "i", "o", "u", "xr", "yt"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(input, prefix) {
			return input + "ay", true
		}
	}
	return input, false
}

func rule2(input string) (string, bool) {
	pos := consonantCluster(input)
	if pos <= 0 {
		return input, false
	}
	return input[pos:] + input[0:pos] + "ay", true
}

func rule3(input string) (string, bool) {
	pos := strings.Index(input, "qu")
	if pos < 0 {
		return input, false
	}
	return input[pos+2:] + input[:pos+2] + "ay", true
}

func rule4(input string) (string, bool) {
	pos := strings.Index(input, "y")
	if pos <= 0 {
		return input, false
	}
	return input[pos:] + input[:pos] + "ay", true
}

func consonantCluster(input string) int {
	for i, c := range input {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			return i
		}
	}
	return -1
}
