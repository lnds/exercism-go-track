package ocr

import (
	"strings"
)

func Recognize(input string) []string {
	lines := strings.Split(input, "\n")
	lines = lines[1:]
	if len(lines)%4 != 0 {
		return nil
	}
	output := []string{}
	for l := 0; l < len(lines); l = l + 4 {
		output = append(output, recognizeNumber(lines[l:l+4]))
	}
	return output
}

func recognizeNumber(lines []string) string {
	digits := []string{}
	for i := 0; i < len(lines[0]); i = i + 3 {
		digits = append(digits, lines[0][i:i+3]+lines[1][i:i+3]+lines[2][i:i+3])
	}
	result := ""
	for _, digit := range digits {
		result += recognizeDigit(digit)
	}
	return result
}

func recognizeDigit(num string) string {
	switch num {
	case " _ | ||_|":
		return "0"
	case "     |  |":
		return "1"
	case "   |_|  |":
		return "4"
	case " _  _||_ ":
		return "2"
	case " _  _| _|":
		return "3"
	case " _ |_  _|":
		return "5"
	case " _ |_ |_|":
		return "6"
	case " _   |  |":
		return "7"
	case " _ |_||_|":
		return "8"
	case " _ |_| _|":
		return "9"
	default:
		return "?"
	}
}
