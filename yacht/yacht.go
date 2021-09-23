package yacht

func Score(dice []int, category string) int {
	switch category {
	case "ones":
		return sum(dice, 1)
	case "twos":
		return sum(dice, 2)
	case "threes":
		return sum(dice, 3)
	case "fours":
		return sum(dice, 4)
	case "fives":
		return sum(dice, 5)
	case "sixes":
		return sum(dice, 6)
	case "full house":
		return fullhouse(dice)
	case "four of a kind":
		return fourofakind(dice)
	case "little straight":
		return littlestraight(dice)
	case "big straight":
		return bigtraight(dice)
	case "choice":
		return choice(dice)
	case "yacht":
		return yacht(dice)
	}
	return 0
}

func yacht(dice []int) int {
	v := dice[0]
	for _, d := range dice[1:] {
		if v != d {
			return 0
		}
	}
	return 50
}

func choice(dice []int) int {
	acum := 0
	for _, d := range dice {
		acum += d
	}
	return acum
}

func fullhouse(dice []int) int {
	counter := map[int]int{}
	for _, d := range dice {
		counter[d]++
	}
	if len(counter) != 2 {
		return 0
	}
	acum := 0
	for k, v := range counter {
		if v < 2 || v > 3 {
			return 0
		}
		acum += k * v
	}
	return acum
}

func fourofakind(dice []int) int {
	counter := map[int]int{}
	for _, d := range dice {
		counter[d]++
	}
	if len(counter) > 2 {
		return 0
	}
	acum := 0
	for k, v := range counter {
		if v != 1 && v != 4 && v != 5 {
			return 0
		}
		if v == 4 {
			acum += k * v
		}
		if v == 5 {
			acum = k * (v - 1)
		}
	}
	return acum
}

func littlestraight(dice []int) int {
	counter := map[int]int{}
	for _, d := range dice {
		counter[d]++
	}
	if len(counter) != 5 {
		return 0
	}
	for d := 1; d < 6; d++ {
		if _, ok := counter[d]; !ok {
			return 0
		}
	}
	return 30
}

func bigtraight(dice []int) int {
	counter := map[int]int{}
	for _, d := range dice {
		counter[d]++
	}
	if len(counter) != 5 {
		return 0
	}
	for d := 2; d < 7; d++ {
		if _, ok := counter[d]; !ok {
			return 0
		}
	}
	return 30
}
func sum(dice []int, n int) int {
	acum := 0
	for _, d := range dice {
		if d == n {
			acum += n
		}
	}
	return acum
}
