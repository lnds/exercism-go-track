package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (result int, err error) {
	if n <= 0 {
		return 0, fmt.Errorf("invalid")
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		result++
	}
	return result, nil
}
