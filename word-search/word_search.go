package wordsearch

import (
	"fmt"
)

type Board struct {
	w, h   int
	matrix [][]rune
}

type Location [2][2]int

func Solve(words, puzzle []string) (result map[string][2][2]int, err error) {
	result = map[string][2][2]int{}
	if len(words) == 0 {
		return result, fmt.Errorf("no words")
	}
	if len(puzzle) == 0 {
		return result, fmt.Errorf("no puzzle")
	}
	board := newBoard(puzzle)
	for _, word := range words {
		if pos, found := board.find(word); found {
			result[word] = pos
		} else {
			err = fmt.Errorf("not found")
		}
	}
	if len(result) == 0 {
		err = fmt.Errorf("not found")
	}
	return
}

func newBoard(puzzle []string) Board {
	w, h := len(puzzle[0]), len(puzzle)
	matrix := [][]rune{}
	for _, p := range puzzle {
		matrix = append(matrix, []rune(p))
	}
	return Board{w: w, h: h, matrix: matrix}
}

func (b Board) find(word string) (Location, bool) {
	delta := len(word) - 1
	if delta >= 0 {
		for y := 0; y < b.h; y++ {
			for x := 0; x < b.w; x++ {
				if b.l2r(x, y, word) {
					return Location{{x, y}, {x + delta, y}}, true
				} else if b.r2l(x, y, delta, word) {
					return Location{{x + delta, y}, {x, y}}, true
				} else if b.t2b(x, y, word) {
					return Location{{x, y}, {x, y + delta}}, true
				} else if b.b2t(x, y, delta, word) {
					return Location{{x, y + delta}, {x, y}}, true
				} else if b.tl2br(x, y, word) {
					return Location{{x, y}, {x + delta, y + delta}}, true
				} else if b.tr2bl(x, y, delta, word) {
					return Location{{x + delta, y}, {x, y + delta}}, true
				} else if b.bl2tr(x, y, delta, word) {
					return Location{{x, y + delta}, {x + delta, y}}, true
				} else if b.br2tl(x, y, delta, word) {
					return Location{{x + delta, y + delta}, {x, y}}, true
				}
			}
		}
	}
	return Location{}, false
}

func (b Board) l2r(x, y int, word string) bool {
	for i, r := range word {
		if b.getChar(x+i, y) != r {
			return false
		}
	}
	return true
}

func (b Board) r2l(x, y, l int, word string) bool {
	for i, r := range word {
		if b.getChar(x+l-i, y) != r {
			return false
		}
	}
	return true
}

func (b Board) t2b(x, y int, word string) bool {
	for i, r := range word {
		if b.getChar(x, y+i) != r {
			return false
		}
	}
	return true
}

func (b Board) b2t(x, y, l int, word string) bool {
	for i, r := range word {
		if b.getChar(x, y+l-i) != r {
			return false
		}
	}
	return true
}

func (b Board) tl2br(x, y int, word string) bool {
	for i, r := range word {
		if b.getChar(x+i, y+i) != r {
			return false
		}
	}
	return true
}

func (b Board) tr2bl(x, y, l int, word string) bool {
	for i, r := range word {
		if b.getChar(x+l-i, y+i) != r {
			return false
		}
	}
	return true
}

func (b Board) bl2tr(x, y, l int, word string) bool {
	for i, r := range word {
		if b.getChar(x+i, y+l-i) != r {
			return false
		}
	}
	return true
}

func (b Board) br2tl(x, y, l int, word string) bool {
	for i, r := range word {
		if b.getChar(x+l-i, y+l-i) != r {
			return false
		}
	}
	return true
}

func (b Board) getChar(x, y int) rune {
	if y < b.h && x < b.w {
		return b.matrix[y][x]
	}
	return rune(0)
}
