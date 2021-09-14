package diamond

import (
	"fmt"
	"strings"
)

func Gen(letter byte) (string, error) {
	if letter < 'A' || letter > 'Z' {
		return "", fmt.Errorf("invalid letter")
	}
	ini := 'A'
	end := rune(letter + 1)
	width := 2*(int(letter)-int(ini)) + 1
	result := []string{}
	for c := ini; c < end; c++ {
		result = append(result, formatChar(c, int(c)-int(ini), width))
	}
	for c := rune(letter - 1); c >= ini; c-- {
		result = append(result, formatChar(c, int(c)-int(ini), width))
	}
	return strings.Join(result, ""), nil
}

func formatChar(c rune, i int, width int) string {
	w := width / 2
	s := string(c)
	if i > 0 {
		s = string(c) + strings.Repeat(" ", 2*i-1) + string(c)
		w = (width - len(s)) / 2
	}
	pad := strings.Repeat(" ", w)
	return fmt.Sprintf("%s%s%s\n", pad, s, pad)
}
