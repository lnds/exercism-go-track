package connect

import "fmt"

func ResultOf(board []string) (string, error) {
	if len(board) == 0 {
		return "", fmt.Errorf("empty board")
	}
	lastCol := len(board[0]) - 1
	lastRow := len(board) - 1

	for i := range board {
		if board[i][0] == 'X' && connected(board, []Pos{}, Pos{0, i}, lastCol, lastRow, 'X') {
			return "X", nil
		}
	}
	for j := range board[0] {
		if board[0][j] == 'O' && connected(board, []Pos{}, Pos{j, 0}, lastCol, lastRow, 'O') {
			return "O", nil
		}
	}
	return "", nil
}

type Pos struct{ x, y int }

func connected(board []string, visited []Pos, pos Pos, lastCol, lastRow int, stone byte) bool {
	if stone == 'X' && pos.x == lastCol {
		return true
	}
	if stone == 'O' && pos.y == lastRow {
		return true
	}
	visited = append(visited, pos)
	deltas := []Pos{{0, 1}, {-1, 1}, {1, -1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range deltas {
		next := Pos{pos.x + d.x, pos.y + d.y}
		if next.y < 0 || next.y > lastRow || next.x < 0 || next.x > lastCol || board[next.y][next.x] != stone || isVisited(next, visited) {
			continue
		}
		if connected(board, visited, next, lastCol, lastRow, stone) {
			return true
		}
	}
	return false
}

func isVisited(pos Pos, visited []Pos) bool {
	for _, v := range visited {
		if v.x == pos.x && v.y == pos.y {
			return true
		}
	}
	return false
}
