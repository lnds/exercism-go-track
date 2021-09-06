package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	triplets := []Triplet{}
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a < b && b < c && a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}

	return triplets
}

func Sum(p int) []Triplet {
	triplets := []Triplet{}
	for _, triplet := range Range(1, p) {
		if triplet[0]+triplet[1]+triplet[2] == p {
			triplets = append(triplets, triplet)
		}
	}
	return triplets
}
