package lsproduct

import (
	"fmt"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	max := -1
	if span < 0 {
		return max, fmt.Errorf("span must not be negative")
	}
	if span > len(digits) {
		return max, fmt.Errorf("span can't be larger that len of digits")
	}
	for i := 0; i < len(digits)-span+1; i++ {
		prod, err := product(digits[i : i+span])
		if err != nil {
			return max, err
		}
		if prod > max {
			max = prod
		}
	}
	return max, nil
}

func product(digits string) (int, error) {
	product := 1
	for _, c := range digits {
		if !unicode.IsDigit(c) {
			return product, fmt.Errorf("invalid digit")
		}
		product *= int(c - '0')
	}
	return product, nil
}
