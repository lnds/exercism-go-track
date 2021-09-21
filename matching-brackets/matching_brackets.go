package brackets

func Bracket(expression string) bool {
	stack := []rune{}
	braces := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	for _, c := range expression {
		switch c {
		case '(', '[', '{':
			stack = append([]rune{c}, stack...)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[0]
			stack = stack[1:]
			if b, ok := braces[top]; ok && b != c {
				return false
			}
		}
	}
	return len(stack) == 0
}
