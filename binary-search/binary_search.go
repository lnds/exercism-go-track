package binarysearch

func SearchInts(array []int, value int) int {
	r := len(array) - 1
	l := 0
	for l <= r {
		m := l + (r-l)/2
		if array[m] < value {
			l = m + 1
		} else if array[m] > value {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
