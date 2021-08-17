// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = -1   // not a triangle
	Equ = iota // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		k = NaT
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		k = NaT
	case a <= 0 || b <= 0 || c <= 0:
		k = NaT
	case a+b < c || a+c < b || b+c < a:
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}
