package sublist

type Relation string

func Sublist(a, b []int) Relation {
	switch {
	case len(a) == len(b) && isEqual(a, b):
		return "equal"
	case len(a) < len(b) && isSubList(a, b):
		return "sublist"
	case len(a) > len(b) && isSubList(b, a):
		return "superlist"
	default:
		return "unequal"
	}
}

func isSubList(shorter, larger []int) bool {
	size := len(shorter)
	top := len(larger) - size + 1
	for i := 0; i < top; i++ {
		if isEqual(shorter, larger[i:i+size]) {
			return true
		}
	}
	return false
}

func isEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
