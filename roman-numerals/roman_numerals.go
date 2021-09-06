package romannumerals

import "fmt"

func part(num, digit int) int {
	return (num % (10 * digit)) / digit
}

func unit(part int, base, mid, limit string) string {
	switch part {
	case 1:
		return base
	case 2:
		return base + base
	case 3:
		return base + base + base
	case 4:
		return base + mid
	case 5:
		return mid
	case 6:
		return mid + base
	case 7:
		return mid + base + base
	case 8:
		return mid + base + base + base
	case 9:
		return base + limit
	default:
		return ""
	}
}

func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number > 3000 {
		return "", fmt.Errorf("invalid number")
	}
	a := unit(part(number, 1000), "M", "", "")
	b := unit(part(number, 100), "C", "D", "M")
	c := unit(part(number, 10), "X", "L", "C")
	d := unit(part(number, 1), "I", "V", "X")
	return a + b + c + d, nil
}
