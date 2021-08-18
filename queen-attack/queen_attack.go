package queenattack

import (
	"fmt"
)

func CanQueenAttack(white, black string) (bool, error) {
	if white == black {
		return false, fmt.Errorf("in same position")
	}
	wr, wc, err := pos(white)
	if err != nil {
		return false, err

	}
	br, bc, err := pos(black)
	if err != nil {
		return false, err

	}
	dr := wr - br
	dc := wc - bc
	result := dr == 0 || dc == 0 || abs(dr) == abs(dc)
	return result, nil
}

func pos(pos string) (int, int, error) {
	if len(pos) != 2 {
		return 0, 0, fmt.Errorf("invalid position")
	}
	row := int(pos[0] - 'a')
	col := int(pos[1] - '1')
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return 0, 0, fmt.Errorf("invalid position")
	}
	return row, col, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
