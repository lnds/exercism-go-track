package prime

func Factors(n int64) []int64 {
	factors := []int64{}
	if n <= 1 {
		return factors
	}
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			break
		}
	}
	if len(factors) == 0 {
		return append(factors, n)
	}
	return append(factors, Factors(n/factors[0])...)
}
