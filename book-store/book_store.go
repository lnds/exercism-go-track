package bookstore

import "sort"

func Cost(basket []int) int {
	mb := map[int]int{}
	for _, v := range basket {
		mb[v]++
	}
	counts := []int{}
	for _, c := range mb {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	if len(counts) == 5 && counts[2] != counts[1] {
		n := min(min(counts[2]/2, counts[1]), counts[0])
		return 2 * n * calcPrice(4)
	}
	price := 0
	used := 0
	for i, p := range counts {
		n := p - used
		price += calcPrice(len(counts)-i) * n
		used += n
	}
	return price
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

const UnitPrice = 800

var Discounts [6]int = [6]int{0, 0, 5, 10, 20, 25}

func calcPrice(count int) int {
	price := count * UnitPrice
	return price * (100 - Discounts[count]) / 100
}
