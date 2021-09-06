package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
	data := []int{}
	for _, c := range isbn {
		if unicode.IsDigit(c) {
			data = append(data, int(c-'0'))
		} else if c == 'x' || c == 'X' {
			if len(data) != 9 {
				return false
			}
			data = append(data, 10)
		}
	}
	if len(data) != 10 {
		return false
	}
	sum := 0
	for i, n := range data {
		sum += (10 - i) * n
	}
	return (sum % 11) == 0
}
