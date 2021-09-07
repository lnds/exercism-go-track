package armstrong

func IsNumber(number int) bool {
	if number == 0 {
		return true
	}

	digits := []int{}
	num := number
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	exp := len(digits)
	sum := 0
	for _, d := range digits {
		sum += pow(d, exp)
	}
	return sum == number
}

func pow(base, exponent int) int {
	result := 1
	for exponent > 0 {
		result *= base
		exponent--
	}
	return result
}
