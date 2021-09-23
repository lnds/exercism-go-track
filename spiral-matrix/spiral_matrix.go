package spiralmatrix

func SpiralMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	k, l := 0, 0
	m := n
	val := 1
	for k < m && l < n {
		for i := l; i < n; i++ {
			result[k][i] = val
			val++
		}
		k++
		for r := k; r < m; r++ {
			result[r][n-1] = val
			val++
		}
		n--
		if k < m {
			for i := n - 1; i >= l; i-- {
				result[m-1][i] = val
				val++
			}
			m--
		}
		if l < n {
			for i := m - 1; i >= k; i-- {
				result[i][l] = val
				val++
			}
			l++
		}

	}
	return result
}
