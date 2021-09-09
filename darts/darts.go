package darts

func Score(x, y float64) int {
	radius := x*x + y*y
	switch {
	case radius <= 1.0:
		return 10
	case radius <= 25.0:
		return 5
	case radius <= 100.0:
		return 1
	default:
		return 0
	}
}
