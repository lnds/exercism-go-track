package dominoes

type Domino [2]int

func MakeChain(dominoes []Domino) ([]Domino, bool) {
	if len(dominoes) == 0 {
		return []Domino{}, true
	}
	bag := []Domino{}
	bag = append(bag, dominoes...)
	stones := []Domino{bag[0]}
	bag = bag[1:]
	return fill(stones, bag)
}

func fill(stones, bag []Domino) ([]Domino, bool) {
	currentEnd := stones[len(stones)-1][1]

	if len(bag) == 0 && stones[0][0] == currentEnd {
		return stones, true
	}

	for i, stone := range bag {
		if p, ok := matchStone(currentEnd, stone); ok {
			newStones := []Domino{}
			newStones = append(newStones, stones...)
			newStones = append(newStones, p)

			remain := []Domino{}
			remain = append(remain, bag[0:i]...)
			remain = append(remain, bag[i+1:]...)
			if fill, ok := fill(newStones, remain); ok {
				return fill, true
			}
		}
	}
	return stones, false
}

func matchStone(end int, stone Domino) (Domino, bool) {
	if stone[0] == end {
		return stone, true
	} else if stone[1] == end {
		return Domino{stone[1], stone[0]}, true
	} else {
		return stone, false
	}
}
