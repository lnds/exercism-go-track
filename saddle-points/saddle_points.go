package matrix

type Pair [2]int

func (m Matrix) Saddle() []Pair {
	rows := m.Rows()
	cols := m.Cols()
	result := []Pair{}
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(cols); c++ {
			n := m[r][c]
			saddle := true
			for _, vi := range rows[r] {
				if n < vi {
					saddle = false
					break
				}
			}
			for _, vj := range cols[c] {
				if n > vj {
					saddle = false
					break
				}
			}
			if saddle {
				result = append(result, Pair{r, c})
			}
		}
	}
	return result
}
