package raindrops

import "strconv"

// Convert a number to raindrop sounds
func Convert(num int) (out string) {
	if num%3 == 0 {
		out += "Pling"
	}
	if num%5 == 0 {
		out += "Plang"
	}
	if num%7 == 0 {
		out += "Plong"
	}
	if out == "" {
		out += strconv.Itoa(num)
	}
	return out
}
