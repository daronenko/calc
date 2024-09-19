package token

import (
	"fmt"
	"strings"
	"unicode"
)

type Token interface{}

func New(value string) (Token, error) {
	var token Token
	var err error

	token, err = NewOperand(value)
	if err == nil {
		return token, nil
	}

	token, err = NewOperator(value)
	if err == nil {
		return token, nil
	}

	token, err = NewBracket(value)
	if err == nil {
		return token, nil
	}

	return nil, fmt.Errorf("unsupported token: %v", err)
}

func Tokenize(expression string) ([]Token, error) {
	var tokens []Token
	var tokenBuilder strings.Builder

	for _, char := range expression {
		switch {
		case unicode.IsSpace(char):
			if tokenBuilder.Len() > 0 {
				// ingore error, because `tokenBuilder` has only digits
				token, _ := New(tokenBuilder.String())

				tokens = append(tokens, token)
				tokenBuilder.Reset()
			}

		case unicode.IsDigit(char) || char == '.':
			tokenBuilder.WriteRune(char)

		case isOperator(string(char)) || isBracket(string(char)):
			if tokenBuilder.Len() > 0 {
				// ingore error, because `tokenBuilder` has only digits
				token, _ := New(tokenBuilder.String())

				tokens = append(tokens, token)
				tokenBuilder.Reset()
			}

			// ingore error, because `char` is operator or bracket
			token, _ := New(string(char))

			tokens = append(tokens, token)

		default:
			return nil, fmt.Errorf("unsupported character: %c", char)
		}
	}

	if tokenBuilder.Len() > 0 {
		token, err := New(tokenBuilder.String())
		if err != nil {
			return nil, fmt.Errorf("cannot tokenize expression: %v", err)
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
