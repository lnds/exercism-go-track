package anagram

import (
	"sort"
	"strings"
)

func Detect(word string, words []string) (candidates []string) {
	norm := normalize(word)
	for _, w := range words {
		if norm == normalize(w) && strings.ToLower(word) != strings.ToLower(w) {
			candidates = append(candidates, w)
		}
	}
	return
}

func normalize(word string) string {
	runes := []rune{}
	for _, r := range strings.ToLower(word) {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
