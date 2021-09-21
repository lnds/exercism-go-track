package transpose

func Transpose(matrix []string) []string {
	width := maxLen(matrix)
	transpose := make([]string, width)

	for i, row := range matrix {
		for j, cell := range row {
			transpose[j] += string(cell)
		}
		rem := maxLen(matrix[i+1:])
		for j := len(row); j < rem; j++ {
			transpose[j] += " "
		}
	}

	return transpose
}

func maxLen(matrix []string) int {
	length := 0
	for _, row := range matrix {
		if len(row) > length {
			length = len(row)
		}
	}
	return length
}
