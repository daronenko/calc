package calculator

import (
	"fmt"

	"github.com/daronenko/calc/internal/notation"
	"github.com/daronenko/calc/internal/token"
	"github.com/daronenko/calc/pkg/stack"
)

func Eval(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0, fmt.Errorf("expression cannot be empty")
	}

	tokens, err := token.Tokenize(expression)
	if err != nil {
		return 0, fmt.Errorf("cannot tokenize expression: %v", err)
	}

	postfixTokens, err := notation.ToPostfix(tokens)
	if err != nil {
		return 0, fmt.Errorf("cannot convert expression to postfix notation: %v", err)
	}

	result, err := evalImpl(postfixTokens)
	if err != nil {
		return 0, fmt.Errorf("cannot evaluate expression: %v", err)
	}

	return result, nil
}

func evalImpl(postfixTokens []token.Token) (float64, error) {
	stack := stack.New[token.Token]()

	for _, t := range postfixTokens {
		switch t := t.(type) {
		case *token.Operand:
			stack.Push(t)

		default:
			rhs, ok := stack.Pop()
			if !ok {
				return 0, fmt.Errorf("no right hand side operand")
			}

			lhs, ok := stack.Pop()
			if !ok {
				return 0, fmt.Errorf("no left hand side operand")
			}

			res, err := t.(*token.Operator).Call(lhs.(*token.Operand), rhs.(*token.Operand))
			if err != nil {
				return 0, fmt.Errorf("error during operation execution: %v", err)
			}

			stack.Push(res)
		}
	}

	// ignore error, because stack always contains at least one element
	result, _ := stack.Pop()

	return result.(*token.Operand).Value(), nil
}
