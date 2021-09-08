package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

func Number(num string) (string, error) {
	result := ""
	for _, c := range num {
		if unicode.IsDigit(c) {
			result += string(c)
		}
	}
	if len(result) < 10 {
		return "", fmt.Errorf("invalid when 9 digits")
	}
	if len(result) > 11 {
		return "", fmt.Errorf("invalid when mor than 11 digits")
	}
	if len(result) == 11 {
		if !strings.HasPrefix(result, "1") {
			return "", fmt.Errorf("invalid when 11 digits does not start with a 1")
		}
		result = result[1:]
	}
	if len(result) == 10 {
		if strings.HasPrefix(result, "0") || strings.HasPrefix(result, "1") {
			return "", fmt.Errorf("invalid if area code starts with 0 or 1")
		}
		if strings.HasPrefix(result[3:], "0") || strings.HasPrefix(result[3:], "1") {
			return "", fmt.Errorf("invalid if exchange starts with 0 or 1")
		}
	}
	return result, nil
}

func AreaCode(num string) (string, error) {
	result, err := Number(num)
	if err != nil {
		return "", err
	}
	return result[0:3], nil
}

func Format(number string) (string, error) {
	num, err := Number(number)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[0:3], num[3:6], num[6:]), nil
}
