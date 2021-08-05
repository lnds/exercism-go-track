package grains

import "errors"

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Square must be between 1 and 64")
	}
	return 1 << (n - 1), nil
}

func Total() (total uint64) {
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		total += square
	}
	return total
}
