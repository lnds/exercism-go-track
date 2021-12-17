package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if t, ok := fnb.(FancyNumber); ok {
		v, _ := strconv.Atoi(t.Value())
		return v
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	n := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(n))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	if t, ok := i.(int); ok {
		return DescribeNumber(float64(t))
	}
	if t, ok := i.(float64); ok {
		return DescribeNumber(t)
	}
	if t, ok := i.(NumberBox); ok {
		return DescribeNumberBox(t)
	}
	if t, ok := i.(FancyNumberBox); ok {
		return DescribeFancyNumberBox(t)
	}
	return "Return to sender"
}
