package wordy

import (
	"strconv"
	"strings"
)

func Answer(expression string) (int, bool) {
	expression = strings.ToLower(expression)
	if !strings.HasSuffix(expression, "?") {
		return 0, false
	}
	return parseQuestion(strings.TrimRight(expression, "?"))
}

func parseQuestion(expression string) (int, bool) {
	tokens := scan(expression)
	if len(tokens) == 0 {
		return 0, false
	}
	if tokens[0].id != What {
		return 0, false
	}
	return parse(tokens[1:])
}

type Token struct {
	id    int
	value int
}

const (
	What = iota
	Plus
	Minus
	Div
	Mult
	Raised
	Cubed
	Number
	Error
)

func scan(expression string) []Token {
	parts := strings.Split(expression, " ")
	tokens := []Token{}
	for _, p := range parts {
		switch p {
		case "what":
			tokens = append(tokens, Token{What, 0})
		case "is", "by", "to", "the", "power":
			continue
		case "plus":
			tokens = append(tokens, Token{Plus, 0})
		case "minus":
			tokens = append(tokens, Token{Minus, 0})
		case "divided":
			tokens = append(tokens, Token{Div, 0})
		case "multiplied":
			tokens = append(tokens, Token{Mult, 0})
		case "raised":
			tokens = append(tokens, Token{Raised, 0})
		default:
			snum := p
			if strings.HasSuffix(p, "th") || strings.HasSuffix(p, "nd") || strings.HasSuffix(p, "st") {
				snum = p[:len(p)-2]
			}
			n, err := strconv.Atoi(snum)
			if err != nil {
				tokens = append(tokens, Token{Error, 0})
			} else {
				tokens = append(tokens, Token{Number, n})
			}
		}
	}
	return tokens
}

func parse(tokens []Token) (int, bool) {
	if len(tokens) == 0 {
		return 0, false
	}
	if tokens[0].id != Number {
		return 0, false
	}
	num := tokens[0].value
	for i := 1; i < len(tokens); i++ {
		tok := tokens[i]
		i++
		if i == len(tokens) {
			return num, false
		}
		switch tok.id {
		case Plus:
			num += tokens[i].value
		case Minus:
			num -= tokens[i].value
		case Mult:
			num *= tokens[i].value
		case Div:
			num /= tokens[i].value
		case Raised:
			num = pow(num, tokens[i].value)
		case Cubed:
			num = pow(num, 3)
		default:
			return num, false
		}
	}
	return num, true
}

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}
