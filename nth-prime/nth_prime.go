package prime

import "math"

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	count := 0
	candidate := 0
	for count < n {
		candidate++
		if isPrime(candidate) {
			count++
		}
	}
	return candidate, true
}

func isPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n < 4:
		return true
	case n%2 == 0:
		return false
	case n < 9:
		return true
	case n%3 == 0:
		return false
	default:
		root := int(math.Sqrt(float64(n)))
		for f := 5; f <= root; f += 6 {
			if n%f == 0 {
				return false
			}
			if n%(f+2) == 0 {
				return false
			}
		}
		return true
	}
}
