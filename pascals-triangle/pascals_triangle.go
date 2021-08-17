package pascal

func Triangle(n int) [][]int {
	triangle := make([][]int, n)
	if n == 0 {
		return triangle
	}

	triangle[0] = []int{1}
	for r := 1; r < n; r++ {
		triangle[r] = row(r, triangle[r-1])
	}
	return triangle
}

func row(n int, prev []int) []int {
	result := make([]int, n+1)
	result[0], result[n] = 1, 1
	for k := 1; k < n; k++ {
		result[k] = prev[k-1] + prev[k]
	}
	return result
}
