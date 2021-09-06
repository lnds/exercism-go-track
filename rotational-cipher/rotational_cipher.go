package rotationalcipher

func RotationalCipher(input string, shift int) string {
	result := ""
	rot := shift % 26
	for _, c := range input {
		ch := c
		if c >= 'a' && c <= 'z' {
			ch = rune((rot+int(c-'a'))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			ch = rune((rot+int(c-'A'))%26 + 'A')
		}
		result += string(ch)
	}
	return result
}
