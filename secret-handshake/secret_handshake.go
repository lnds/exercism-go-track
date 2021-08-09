package secret

func Handshake(code uint) (result []string) {
	if code&0x01 == 0x01 {
		result = append(result, "wink")
	}
	if code&0x02 == 0x02 {
		result = append(result, "double blink")
	}
	if code&0x04 == 0x04 {
		result = append(result, "close your eyes")
	}
	if code&0x08 == 0x08 {
		result = append(result, "jump")
	}
	if code&0x10 == 0x10 {
		reverse(result[:])
	}
	return result
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
