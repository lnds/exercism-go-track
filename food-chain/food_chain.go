package foodchain

import (
	"fmt"
	"strings"
)

var animals = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var scondLine = []string{
	"",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
	"She's dead, of course!",
}

func Verse(i int) string {
	stansas := []string{}
	animal := animals[i-1]
	stansas = append(stansas, fmt.Sprintf("I know an old lady who swallowed a %s.", animal))
	if i == 8 {
		stansas = append(stansas, scondLine[i-1])
	} else if i < len(animals) {
		if i > 1 {
			stansas = append(stansas, scondLine[i-1])
			for ; i > 1; i-- {
				animal := animals[i-1]
				previous := animals[i-2]
				if previous == "spider" {
					previous += " that wriggled and jiggled and tickled inside her"
				}
				stansas = append(stansas, fmt.Sprintf("She swallowed the %s to catch the %s.", animal, previous))
			}
		}
		stansas = append(stansas, "I don't know why she swallowed the fly. Perhaps she'll die.")
	}
	return strings.Join(stansas, "\n")
}

func Verses(i, j int) string {
	verses := []string{}
	for k := i; k <= j; k++ {
		verses = append(verses, Verse(k))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
