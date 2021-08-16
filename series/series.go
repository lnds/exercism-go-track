package series

func All(n int, s string) (result []string) {
	for i := 0; i < len(s)+1-n; i++ {
		result = append(result, UnsafeFirst(n, s[i:]))
	}
	return
}

func UnsafeFirst(n int, s string) (result string) {
	for i := 0; i < n; i++ {
		result += string(s[i])
	}
	return
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return UnsafeFirst(n, s), true
}
