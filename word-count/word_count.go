package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(text string) Frequency {
	result := make(map[string]int)
	splitter := func(c rune) bool {
		return c != '\'' && (unicode.IsSpace(c) || unicode.IsPunct(c) || unicode.IsSymbol(c))
	}
	words := strings.FieldsFunc(strings.ToLower(text), splitter)
	for _, word := range words {
		word := strings.Trim(word, "'")
		result[word]++
	}
	return result
}
