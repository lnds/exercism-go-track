package summultiples

// SumMultiples of any factor in factors
func SumMultiples(limit int, factors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, f := range factors {
			if f > 0 && i%f == 0 {
				sum += i
				break
			}
		}
	}
	return
}
