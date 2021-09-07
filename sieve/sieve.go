package sieve

func Sieve(limit int) []int {
	sieve := make([]bool, limit+1)
	for i := 0; i <= limit; i++ {
		sieve[i] = true
	}
	for i := 2; i <= limit; i++ {
		for j := i + i; j <= limit; j += i {
			sieve[j] = false
		}
	}
	result := []int{}
	for i := 2; i <= limit; i++ {
		if sieve[i] {
			result = append(result, i)
		}
	}
	return result
}
