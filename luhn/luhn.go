package luhn

func Valid(code string) bool {
	acum := 0
	j := 0
	for i := len(code) - 1; i >= 0; i-- {
		ch := code[i]
		if ch == ' ' {
			continue
		}
		if ch < '0' || ch > '9' {
			return false
		}
		num := int(ch - '0')
		if j%2 == 0 {
			acum += num
		} else if num*2 > 9 {
			acum += num*2 - 9
		} else {
			acum += num * 2
		}
		j++
	}
	return (j > 1) && (acum%10) == 0
}
