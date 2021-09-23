package railfence

import (
	"strings"
)

func Encode(plain string, rails int) string {
	fences := make([]string, rails)
	row, up := 0, true
	for _, c := range plain {
		fences[row] += string(c)
		row, up = zigzag(row, rails, up)
	}
	return strings.Join(fences, "")
}

func Decode(cipher string, rails int) string {
	fencesLen := make([]int, rails)
	row, up := 0, true
	for i := 0; i < len(cipher); i++ {
		fencesLen[row]++
		row, up = zigzag(row, rails, up)
	}
	fences := make([]string, rails)
	i := 0
	for f := range fencesLen {
		n := fencesLen[f]
		fences[f] = cipher[i : i+n]
		i += n
	}
	result := ""
	row, up = 0, true
	for f := 0; f < len(cipher); f++ {
		fence := fences[row]
		result += string(fence[0])
		fences[row] = fence[1:]
		row, up = zigzag(row, rails, up)
	}

	return result

}

func zigzag(row, rails int, up bool) (int, bool) {
	if up {
		if row < rails-1 {
			return row + 1, true
		} else {
			return row - 1, false
		}
	} else {
		if row > 0 {
			return row - 1, false
		} else {
			return row + 1, true
		}
	}
}
