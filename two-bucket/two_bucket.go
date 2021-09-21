package twobucket

import "fmt"

func Solve(capacity1, capacity2, goal int, start string) (goalBucket string, moves int, otherBucket int, err error) {

	if goal > capacity2 || goal%gcd(capacity1, capacity2) != 0 {
		err = fmt.Errorf("not possible")
		return
	}
	if capacity1 == 0 {
		err = fmt.Errorf("invalid first bucket size")
		return
	}
	if goal == 0 {
		err = fmt.Errorf("invalid goal")
		return
	}
	switch {
	case start == "one" && goal == capacity2:
		moves = 2
		goalBucket = "two"
		otherBucket = capacity1
		return
	case start == "one" && goal == capacity1:
		moves = 1
		goalBucket = "one"
		otherBucket = 0
		return
	case start == "two" && goal == capacity1:
		moves = 2
		goalBucket = "one"
		otherBucket = capacity2
		return
	case start == "two" && goal == capacity2:
		moves = 1
		goalBucket = "two"
		otherBucket = 0
		return
	case start == "one":
		from, to, steps := pour(capacity1, capacity2, goal)
		moves = steps
		if goal == from {
			goalBucket = "one"
		} else {
			goalBucket = "two"
		}
		if goal == from {
			otherBucket = to
		} else {
			otherBucket = from
		}
		return
	case start == "two":
		from, to, steps := pour(capacity2, capacity1, goal)
		moves = steps
		if goal == from {
			goalBucket = "two"
		} else {
			goalBucket = "one"
		}
		if goal == from {
			otherBucket = to
		} else {
			otherBucket = from
		}
		return
	default:
		err = fmt.Errorf("invalid bucket name")
		return
	}
}

func pour(capacity1, capacity2, goal int) (int, int, int) {
	from := capacity1
	to := 0
	step := 1
	for from != goal && to != goal {
		temp := min(from, capacity2-to)
		to += temp
		from -= temp
		step++
		if from == goal || to == goal {
			break
		}
		if from == 0 {
			from = capacity1
			step++
		}
		if to == capacity2 {
			to = 0
			step++
		}
	}
	return from, to, step
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
