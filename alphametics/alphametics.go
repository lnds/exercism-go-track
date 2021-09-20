package alphametics

import (
	"fmt"
	"strings"
)

func Solve(expression string) (map[string]int, error) {
	summands, sum, err := parse(expression)
	if err != nil {
		return nil, err
	}

	values := map[rune]int{}
	letters := []rune{}
	signs := strings.Join(summands, "") + sum
	for _, letter := range signs {
		if _, exists := values[letter]; !exists {
			values[letter] = 0
			letters = append(letters, letter)
		}
	}
	if len(letters) > 10 {
		return nil, fmt.Errorf("only 10 letters")
	}

	digits := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// check first combination
	assignValues(letters, digits, values)
	if check(summands, sum, values) {
		return transform(values), nil
	}

	// calc permutations, using Heap algorithm
	i := 0
	c := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i < 10 {
		if c[i] >= i {
			c[i] = 0
			i++
		} else {
			if i%2 == 0 {
				digits[0], digits[i] = digits[i], digits[0]
			} else {
				digits[c[i]], digits[i] = digits[i], digits[c[i]]
			}
			assignValues(letters, digits, values)
			if check(summands, sum, values) {
				return transform(values), nil
			}
			c[i]++
			i = 0
		}
	}
	return nil, fmt.Errorf("not found")
}

// transform to expected result format
func transform(values map[rune]int) map[string]int {
	result := map[string]int{}
	for k, v := range values {
		result[string(k)] = v
	}
	return result
}

// asigns guessed digits to map of values
func assignValues(letters []rune, digits [10]int, values map[rune]int) {
	i := 0
	for _, letter := range letters {
		values[letter] = digits[i]
		i++
	}
}

// parse the input
func parse(input string) (sumands []string, sum string, err error) {
	parts := strings.Split(input, "==")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid format")
		return
	}
	sum = strings.TrimSpace(parts[1])
	sumands = strings.Split(parts[0], "+")
	for i, s := range sumands {
		sumands[i] = strings.TrimSpace(s)
	}
	return
}

// check equation, using map of runes->int
func check(summands []string, sum string, digits map[rune]int) bool {
	left := 0
	for _, s := range summands {
		v := value(s, digits)
		if v == 0 {
			return false
		}
		left += v
	}
	right := value(sum, digits)
	return left == right
}

// convert a string to a number
func value(num string, digits map[rune]int) int {
	if num == "" {
		return 0
	}
	acc := 0
	for i, c := range num {
		digit, ok := digits[c]
		if !ok || (i == 0 && digit == 0) {
			return 0
		}
		acc = acc*10 + digit
	}
	return acc
}
