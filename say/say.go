package say

func Say(num int64) (string, bool) {
	digits := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	tens := []string{
		"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety", "one hundred",
	}
	switch {
	case num >= 0 && num < 20:
		return digits[num], true
	case num%10 == 0 && num >= 20 && num <= 100:
		return tens[(num-20)/10], true
	case num >= 21 && num <= 99:
		p, _ := Say(num / 10 * 10)
		s, _ := Say(num % 10)
		return p + "-" + s, true
	case num >= 101 && num < 1000:
		return suffix(num, 100, "hundred"), true
	case num >= 1_000 && num < 1_000_000:
		return suffix(num, 1_000, "thousand"), true
	case num >= 1_000_000 && num < 1_000_000_000:
		return suffix(num, 1_000_000, "million"), true
	case num >= 1_000_000_000 && num < 1_000_000_000_000:
		return suffix(num, 1_000_000_000, "billion"), true
	}
	return "", false
}

func suffix(n, m int64, suf string) string {
	a := n / m
	b := n % m
	if b == 0 {
		s, _ := Say(a)
		return s + " " + suf
	} else {
		s, _ := Say(a)
		p, _ := Say(b)
		return s + " " + suf + " " + p
	}
}
