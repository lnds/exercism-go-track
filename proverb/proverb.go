// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) (result []string) {
	if len(rhyme) == 0 {
		return
	}
	elem := rhyme[0]
	for i := 1; i < len(rhyme); i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	result = append(result, fmt.Sprintf("And all for the want of a %s.", elem))
	return result
}
