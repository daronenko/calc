package notation

import (
	"fmt"

	"github.com/daronenko/calc/internal/token"
	"github.com/daronenko/calc/pkg/stack"
)

func ToPostfix(tokens []token.Token) ([]token.Token, error) {
	var postfix []token.Token
	operatorStack := stack.New[token.Token]()

	for _, t := range tokens {
		switch t := t.(type) {
		case *token.Operand:
			postfix = append(postfix, t)

		case *token.Operator:
			for !operatorStack.IsEmpty() {
				top, ok := operatorStack.Top()
				if !ok {
					return nil, fmt.Errorf("cannot get top value of empty stack")
				}
				topOperator, isOperator := (*top).(*token.Operator)

				if !isOperator || topOperator.Precedence() < t.Precedence() {
					break
				}

				// ignore error, because `operatorStack` has top value
				operator, _ := operatorStack.Pop()

				postfix = append(postfix, operator)
			}

			operatorStack.Push(t)

		case *token.Bracket:
			if t.IsOpeningBracket() {
				operatorStack.Push(t)
				continue
			}

			openingBracketFound := false

			for !operatorStack.IsEmpty() {
				top, err := operatorStack.Pop()
				if !err {
					return nil, fmt.Errorf("cannot pop value from empty stack: %v", err)
				}

				topBracket, isBracket := top.(*token.Bracket)
				if isBracket && topBracket.IsOpeningBracket() {
					openingBracketFound = true
					break
				}

				postfix = append(postfix, top)
			}

			if !openingBracketFound {
				return nil, fmt.Errorf("invalid bracket sequence")
			}
		}
	}

	for !operatorStack.IsEmpty() {
		top, _ := operatorStack.Pop()

		if _, isBracket := top.(*token.Bracket); isBracket {
			return nil, fmt.Errorf("expression has invalid bracket sequence")
		}

		postfix = append(postfix, top)
	}

	return postfix, nil
}
