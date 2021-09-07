package cryptosquare

import (
	"fmt"
	"strings"
	"unicode"
)

func Encode(plain string) string {
	cipher := ""
	for _, c := range strings.ToLower(plain) {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			cipher += string(c)
		}
	}
	if cipher == "" {
		return ""
	}

	r, c := size(cipher)
	cipher = fmt.Sprintf("%-*s", r*c, cipher)
	rectangle := rect(cipher, c)
	transpose := make([][]rune, c)
	for i := range transpose {
		transpose[i] = make([]rune, r)
	}
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			transpose[i][j] = rectangle[j][i]
		}
	}

	cipher = ""
	for i, row := range transpose {
		cipher += string(row)
		if i < len(transpose)-1 {
			cipher += " "
		}
	}
	return cipher
}

func size(cipher string) (int, int) {
	lc := len(cipher)
	x := 1
	for x*x < lc && (x+1)*x < lc {
		x++
	}
	if x*x >= lc {
		return x, x
	}
	return x, x + 1
}

func rect(cipher string, c int) [][]rune {
	rectangle := [][]rune{}
	for i := 0; i < len(cipher)-c+1; i += c {
		rectangle = append(rectangle, []rune(cipher[i:i+c]))
	}
	return rectangle
}
