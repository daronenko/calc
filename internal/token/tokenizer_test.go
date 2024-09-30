package token_test

import (
	"reflect"
	"testing"

	"github.com/daronenko/calc/internal/token"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input       string
		expected    []token.Token
		expectError bool
	}{
		{
			input: "3 + 4",
			expected: []token.Token{
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("4")),
			},
		},
		{
			input: "10 - 2 * 5",
			expected: []token.Token{
				ignoreError(token.NewOperand("10")),
				ignoreError(token.NewOperator("-")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewOperand("5")),
			},
		},
		{
			input: "5 + (3 * 2)",
			expected: []token.Token{
				ignoreError(token.NewOperand("5")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewBracket(")")),
			},
		},
		{
			input: "(10 / 5) + 7",
			expected: []token.Token{
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("10")),
				ignoreError(token.NewOperator("/")),
				ignoreError(token.NewOperand("5")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("7")),
			},
		},
		{
			input: "7 + (3 * (10 / 5))",
			expected: []token.Token{
				ignoreError(token.NewOperand("7")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("10")),
				ignoreError(token.NewOperator("/")),
				ignoreError(token.NewOperand("5")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewBracket(")")),
			},
		},
		{
			input:    "",
			expected: []token.Token{},
		},
		{
			input:       "a + 1",
			expected:    nil,
			expectError: true,
		},
		{
			input:       "7 % * 3",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		tokens, err := token.Tokenize(tt.input)

		if hasError := err != nil; hasError != tt.expectError {
			t.Error("expected error")
		}

		if !tt.expectError && !(len(tokens)+len(tt.expected) == 0 || reflect.DeepEqual(tokens, tt.expected)) {
			t.Errorf("got '%v', expected '%v'", tokens, tt.expected)
		}
	}
}

func ignoreError(value interface{}, _ error) interface{} {
	return value
}
