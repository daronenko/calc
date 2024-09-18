package notation_test

import (
	"testing"

	"github.com/daronenko/calc/internal/notation"
	"github.com/daronenko/calc/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestToPostfix(t *testing.T) {
	tests := []struct {
		input       []token.Token
		expected    []token.Token
		expectError bool
	}{
		{
			input: []token.Token{
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("4")),
			},
			expected: []token.Token{
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperand("4")),
				ignoreError(token.NewOperator("+")),
			},
		},
		{
			input: []token.Token{
				ignoreError(token.NewOperand("7")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewOperand("8")),
				ignoreError(token.NewOperator("/")),
				ignoreError(token.NewOperand("9")),
			},
			expected: []token.Token{
				ignoreError(token.NewOperand("7")),
				ignoreError(token.NewOperand("8")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewOperand("9")),
				ignoreError(token.NewOperator("/")),
			},
		},
		{
			input: []token.Token{
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("1")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewOperand("3")),
			},
			expected: []token.Token{
				ignoreError(token.NewOperand("1")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperator("*")),
			},
		},
		{
			input: []token.Token{
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("5")),
				ignoreError(token.NewOperator("-")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("6")),
				ignoreError(token.NewOperator("/")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("3")),
			},
			expected: []token.Token{
				ignoreError(token.NewOperand("5")),
				ignoreError(token.NewOperand("6")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("/")),
				ignoreError(token.NewOperator("-")),
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperator("+")),
			},
		},
		{
			input: []token.Token{
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("4")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewBracket(")")),
			},
			expected: []token.Token{
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperand("4")),
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperator("*")),
			},
		},
		{
			input: []token.Token{
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("4")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("3")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewBracket(")")),
			},
			expectError: true,
		},
		{
			input: []token.Token{
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("4")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("1")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("3")),
			},
			expectError: true,
		},
		{
			input: []token.Token{
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("*")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("4")),
				ignoreError(token.NewOperator("/")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewBracket("(")),
				ignoreError(token.NewOperand("1")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("1")),
				ignoreError(token.NewBracket(")")),
				ignoreError(token.NewBracket(")")),
			},
			expected: []token.Token{
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperand("4")),
				ignoreError(token.NewOperand("1")),
				ignoreError(token.NewOperand("2")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperand("1")),
				ignoreError(token.NewOperator("+")),
				ignoreError(token.NewOperator("/")),
				ignoreError(token.NewOperator("*")),
			},
		},
	}

	for _, tt := range tests {
		t.Run("expression", func(t *testing.T) {
			result, err := notation.ToPostfix(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func ignoreError(value interface{}, _ error) interface{} {
	return value
}
