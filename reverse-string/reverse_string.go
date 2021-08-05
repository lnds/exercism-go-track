package reverse

func Reverse(str string) (out string) {
	for _, ch := range str {
		out = string(ch) + out
	}
	return out
}
