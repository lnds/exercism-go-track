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

	signs := strings.Join(summands, "") + sum

	digits := intSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	values := newValues(signs)
	var generate func(int, intSlice) (map[string]int, error)
	generate = func(n int, guess intSlice) (map[string]int, error) {
		if n == 1 {
			A := guess.Copy()
			i := 0
			for r := range values {
				values[r] = A[i]
				i++
			}
			if check(summands, sum, values) {
				return values, nil
			}
		} else {
			for i := 0; i < n; i++ {
				v, e := generate(n-1, guess)
				if e == nil {
					return v, e
				}
				if n%2 == 0 {
					guess.Swap(i, n-1)
				} else {
					guess.Swap(0, n-1)
				}
			}
		}
		return nil, fmt.Errorf("not found")

	}
	return generate(digits.Len(), digits)
}

func newValues(signs string) map[string]int {
	values := map[string]int{}
	for _, r := range signs {
		values[string(r)] = 0
	}
	return values
}

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

func check(summands []string, sum string, digits map[string]int) bool {
	left := 0
	for _, s := range summands {
		if v, ok := value(s, digits); !ok {
			return false
		} else {
			left += v
		}
	}
	if right, ok := value(sum, digits); !ok {
		return false
	} else {
		return left == right
	}
}

func value(num string, digits map[string]int) (int, bool) {
	if num == "" {
		return 0, false
	}
	acc := 0
	for i, c := range num {
		if digit, ok := digits[string(c)]; !ok {
			return 0, false
		} else {
			if i == 0 && digit == 0 {
				return 0, false
			}
			if i == 0 {
				acc = digit
			} else {
				acc = acc*10 + digit
			}
		}
	}
	return acc, true
}

type intSlice []int

func (p intSlice) Len() int      { return len(p) }
func (p intSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p intSlice) Copy() intSlice {
	A := make(intSlice, p.Len())
	copy(A, p)
	return A
}
