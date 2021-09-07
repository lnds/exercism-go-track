package allyourbase

import "fmt"

func ConvertToBase(fromBase int, digits []int, toBase int) ([]int, error) {
	if fromBase <= 1 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if toBase <= 1 {
		return nil, fmt.Errorf("output base must be >= 2")

	}
	if len(digits) == 0 {
		return []int{0}, nil
	}
	value := 0
	for i := 0; i < len(digits); i++ {
		if digits[i] < 0 || digits[i] >= fromBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		value = value*fromBase + digits[i]
	}

	if value == 0 {
		return []int{0}, nil
	}

	result := []int{}
	for value > 0 {
		result = append([]int{value % toBase}, result...)
		value /= toBase
	}
	return result, nil
}
