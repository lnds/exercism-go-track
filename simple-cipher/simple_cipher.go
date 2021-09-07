package cipher

import (
	"strings"
)

type VignereCipher struct {
	key string
}

func NewVigenere(key string) Cipher {
	if invalidKey(key) {
		return nil
	}
	return VignereCipher{
		key: key,
	}
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(shift int) Cipher {
	if shift <= 0 {
		shift = 26 + shift
	}
	if shift <= 0 || shift >= 26 {
		return nil
	}
	return NewVigenere(string(byte('a' + shift)))
}

func (cipher VignereCipher) Encode(input string) string {
	return Transform(cipher.key, strings.ToLower(input), encodeChar)
}

func (cipher VignereCipher) Decode(input string) string {
	return Transform(cipher.key, input, decodeChar)
}

func Transform(key, text string, f func(byte, byte) (byte, bool)) string {
	key_len := len(key)
	result := ""
	j := 0
	for i := 0; i < len(text); i++ {
		k := key[j%key_len]
		c := text[i]
		if ch, ok := f(c, k); ok {
			j++
			result += string(ch)
		}
	}
	return result
}

func encodeChar(c, key byte) (byte, bool) {
	if c >= 'a' && c <= 'z' {
		return byte('a' + (26+(c-'a')+(key-'a'))%26), true
	}
	return c, false
}

func decodeChar(c, key byte) (byte, bool) {
	if c >= 'a' && c <= 'z' {
		return byte('a' + (26+(c-'a')-(key-'a'))%26), true
	}
	return c, false
}

func invalidKey(key string) bool {
	if key == "" {
		return true
	}
	sum := 0
	for i := 0; i < len(key); i++ {
		ch := key[i]
		if ch < 'a' || ch > 'z' {
			return true
		}
		sum += int(ch - 'a')
	}
	return sum == 0
}
