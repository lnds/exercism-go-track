package beer

import (
	"fmt"
)

func Verse(n int) (string, error) {
	switch {
	case n == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case n == 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case n == 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	case n > 2 && n < 100:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	default:
		return "", fmt.Errorf("invalid verse")
	}
}

func Verses(last, first int) (string, error) {
	verses := ""
	for i := last; i >= first; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses = verses + verse + "\n"
	}
	if verses == "" {
		return verses, fmt.Errorf("invalid order")
	}
	return verses, nil
}

func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}
