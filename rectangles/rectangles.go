package rectangles

func Count(lines []string) int {
	h := len(lines)
	if h == 0 {
		return 0
	}
	w := len(lines[0])
	if w == 0 {
		return 0
	}
	vertices := findVertices(lines)
	count := 0

	for _, p := range vertices {
		for _, q := range vertices {
			if p[0] < q[0] && p[1] < q[1] && connected(p, q, lines) {
				count++
			}
		}
	}
	return count
}

type Point [2]int

func findVertices(lines []string) []Point {
	vertices := []Point{}
	for y, row := range lines {
		for x, c := range row {
			if c == '+' {
				vertices = append(vertices, Point{x, y})
			}
		}
	}
	return vertices
}

func connected(p0, p1 Point, lines []string) bool {
	x0 := p0[0]
	y0 := p0[1]

	x1 := p1[0]
	y1 := p1[1]

	if lines[y0][x1] != '+' || lines[y1][x0] != '+' {
		return false
	}

	for _, c0 := range lines[y0][x0:x1] {
		for _, c1 := range lines[y1][x0:x1] {
			if (c0 != '-' && c0 != '+') || (c1 != '-' && c1 != '+') {
				return false
			}
		}
	}
	for _, row := range lines[y0:y1] {
		c0 := row[x0]
		c1 := row[x1]
		if (c0 != '|' && c0 != '+') || (c1 != '|' && c1 != '+') {
			return false
		}
	}
	return true
}
