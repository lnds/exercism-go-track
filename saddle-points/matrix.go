package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(matrixAsString string) (*Matrix, error) {
	var matrix Matrix
	for i, r := range strings.Split(matrixAsString, "\n") {
		var row []int
		for _, c := range strings.Split(strings.TrimSpace(r), " ") {
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		if i > 0 && len(matrix[0]) != len(row) {
			return nil, fmt.Errorf("non rectangular")
		}
		matrix = append(matrix, row)
	}
	return &matrix, nil
}

func (matrix Matrix) Rows() [][]int {
	n := len(matrix)
	m := len(matrix[0])
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
		for j := range result[i] {
			result[i][j] = matrix[i][j]
		}
	}
	return result
}

func (matrix Matrix) Cols() [][]int {
	n := len(matrix[0])
	m := len(matrix)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
		for j := range result[i] {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

func (matrix Matrix) Set(i, j, value int) bool {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		return false
	}
	matrix[i][j] = value
	return true
}
