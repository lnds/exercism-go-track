package minesweeper

import "fmt"

func (b Board) Count() error {
	w, h, err := validate(b)
	if err != nil {
		return err
	}
	for j := 1; j < h-1; j++ {
		for i := 1; i < w-1; i++ {
			if b[j][i] == ' ' {
				b[j][i] = calc(j, i, h, w, b)
			}
		}
	}
	return nil
}

func calc(j, i, h, w int, b Board) byte {
	mines := mine(j-1, i-1, h, w, b) + mine(j-1, i, h, w, b) + mine(j-1, i+1, h, w, b) + mine(j, i-1, h, w, b) + mine(j, i+1, h, w, b) + mine(j+1, i-1, h, w, b) + mine(j+1, i, h, w, b) + mine(j+1, i+1, h, w, b)
	if mines > 0 {
		return '0' + mines
	}
	return ' '
}

func mine(j, i, h, w int, b Board) byte {
	if j < 1 || i < 1 || j > h-1 || i > w-1 {
		return 0
	}
	if b[j][i] == '*' {
		return 1
	}
	return 0
}

func validate(b Board) (w, h int, err error) {
	if len(b) == 0 {
		err = fmt.Errorf("empty board")
		return
	}
	h = len(b)
	w = len(b[0])

	for j := range b {
		if len(b[j]) != w {
			err = fmt.Errorf("unaligned")
			return
		}
	}

	if b[0][0] != '+' || b[0][w-1] != '+' || b[h-1][0] != '+' || b[h-1][w-1] != '+' {
		err = fmt.Errorf("invalid corners")
		return
	}
	for i := 1; i < w-1; i++ {
		if b[0][i] != '-' || b[h-1][i] != '-' {
			err = fmt.Errorf("invalid border")
			return
		}
	}
	for j := 1; j < h-1; j++ {
		if b[j][0] != '|' || b[j][w-1] != '|' {
			err = fmt.Errorf("invalid border")
			return
		}
	}
	for j := 1; j < h-1; j++ {
		for i := 1; i < w-1; i++ {
			if b[j][i] != ' ' && b[j][i] != '*' {
				err = fmt.Errorf("invalid character")
				return
			}
		}
	}

	return
}
