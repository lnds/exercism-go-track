package scale

import "strings"

var Flat = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

var SharpScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var FlatScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

func Scale(tonic, intervals string) []string {
	scale := SharpScale
	for _, note := range Flat {
		if note == tonic {
			scale = FlatScale
			break
		}
	}

	pos := -1
	for i, note := range scale {
		if strings.EqualFold(note, tonic) {
			pos = i
			break
		}
	}

	// chromatic
	if intervals == "" {
		intervals = "mmmmmmmmmmmm"
	}
	return newScale(intervals, pos, scale)
}

func newScale(intervals string, pos int, scale []string) []string {
	notes := []string{}
	p := pos
	notes = append(notes, scale[p])
	for _, interval := range intervals {
		switch interval {
		case 'M':
			p = (p + 2) % 12
		case 'm':
			p = (p + 1) % 12
		case 'A':
			p = (p + 3) % 12
		}
		notes = append(notes, scale[p])
	}
	return notes[:len(notes)-1]
}
