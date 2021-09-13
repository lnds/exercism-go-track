package forth

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Forth(statements []string) (result []int, err error) {
	env := NewEnvironment()
	for _, statement := range statements {
		err = env.Eval(statement)
		if err != nil {
			return nil, err
		}
	}
	return env.stack, nil
}

type Environment struct {
	stack []int
	defs  map[string][]string
}

func NewEnvironment() *Environment {
	return &Environment{stack: []int{}, defs: make(map[string][]string)}
}

func (env *Environment) Eval(statement string) error {
	return env.Run(strings.ToUpper(statement))
}

func (env *Environment) Run(statements string) error {
	tokens := strings.FieldsFunc(statements, func(c rune) bool {
		return !unicode.IsDigit(c) && !unicode.IsLetter(c) && !unicode.IsSymbol(c) && !unicode.IsPunct(c)
	})
	return env.Parse(tokens)
}

func (env *Environment) Parse(tokens []string) error {
	if len(tokens) == 0 {
		return nil
	}
	if tokens[0] == ":" {
		return env.ParseDef(tokens)
	}

	for _, word := range tokens {
		err := env.ParseWord(word)
		if err != nil {
			return err
		}
	}
	return nil
}

func (env *Environment) ParseDef(tokens []string) error {

	if tokens[len(tokens)-1] != ";" {
		return fmt.Errorf("invalid definition")
	}
	word := tokens[1]

	if _, err := strconv.Atoi(word); err == nil {
		return fmt.Errorf("can't redefine numbers")
	}

	defintion := tokens[2 : len(tokens)-1]
	body := []string{}
	for _, w := range defintion {
		if wdef, ok := env.GetDefinition(w); ok {
			body = append(body, wdef...)
		} else {
			body = append(body, w)
		}
	}
	env.Define(word, body)
	return nil
}

func (env *Environment) ParseWord(token string) (err error) {
	if env.IsDefined(token) {
		return env.ParseValue(token)
	}
	switch token {
	case ":":
		return fmt.Errorf("invalid definition")
	case "+":
		return env.Add()
	case "-":
		return env.Sub()
	case "*":
		return env.Mul()
	case "/":
		return env.Div()
	case "DUP":
		return env.Dup()
	case "DROP":
		return env.Drop()
	case "SWAP":
		return env.Swap()
	case "OVER":
		return env.Over()
	default:
		return env.ParseValue(token)
	}
}

func (env *Environment) ParseValue(token string) error {
	if val, err := strconv.Atoi(token); err == nil {
		env.Push(val)
		return nil
	}
	if tokens, ok := env.GetDefinition(token); !ok {
		return fmt.Errorf("word %s not defined", token)
	} else {
		newEnv := NewEnvironment()
		newEnv.CopyStack(env)
		err := newEnv.Eval(strings.Join(tokens, " "))
		if err == nil {
			env.CopyStack(newEnv)
		}
		return err
	}
}

func (env *Environment) Define(word string, tokens []string) {
	env.defs[word] = tokens
}

func (env *Environment) GetDefinition(word string) ([]string, bool) {
	v, ok := env.defs[word]
	return v, ok
}

func (env *Environment) IsDefined(word string) bool {
	_, ok := env.defs[word]
	return ok
}

func (env *Environment) Op(operation func(int, int) (int, error)) error {
	if len(env.stack) < 2 {
		return fmt.Errorf("empty stack")
	}
	var a, b int
	a, _ = env.Pop()
	b, _ = env.Pop()
	c, err := operation(a, b)
	if err != nil {
		return err
	}
	env.stack = append(env.stack, c)
	return nil
}

func (env *Environment) Add() error {
	return env.Op(func(x, y int) (int, error) { return x + y, nil })
}

func (env *Environment) Sub() error {
	return env.Op(func(x, y int) (int, error) { return y - x, nil })
}

func (env *Environment) Mul() error {
	return env.Op(func(x, y int) (int, error) { return x * y, nil })
}

func (env *Environment) Div() error {
	return env.Op(func(x, y int) (int, error) {
		if x == 0 {
			return 0, fmt.Errorf("div by 0")
		}
		return y / x, nil
	})
}

func (env *Environment) Dup() error {
	if len(env.stack) == 0 {
		return fmt.Errorf("empty stack")
	}
	env.stack = append(env.stack, env.stack[len(env.stack)-1])
	return nil
}

func (env *Environment) Drop() error {
	if len(env.stack) == 0 {
		return fmt.Errorf("empty stack")
	}
	env.stack = env.stack[:len(env.stack)-1]
	return nil
}

func (env *Environment) Swap() error {
	if len(env.stack) < 2 {
		return fmt.Errorf("empty stack")
	}
	var a, b int
	a, _ = env.Pop()
	b, _ = env.Pop()
	env.Push(a)
	env.Push(b)
	return nil
}

func (env *Environment) Over() error {
	if len(env.stack) < 2 {
		return fmt.Errorf("empty stack")
	}
	var a, b int
	a, _ = env.Pop()
	b, _ = env.Pop()
	env.Push(b)
	env.Push(a)
	env.Push(b)
	return nil
}

func (env *Environment) Push(val int) {
	env.stack = append(env.stack, val)
}

func (env *Environment) Pop() (int, error) {
	if len(env.stack) == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	v := env.stack[len(env.stack)-1]
	env.stack = env.stack[:len(env.stack)-1]
	return v, nil
}

func (env *Environment) CopyStack(other *Environment) {
	env.stack = []int{}
	env.stack = append(env.stack, other.stack...)
}
