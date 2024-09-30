package calculator_test

import (
	"testing"

	"github.com/daronenko/calc/internal/calculator"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expression  string
		expected    float64
		expectError bool
	}{
		{
			expression: "3 + 4",
			expected:   7,
		},
		{
			expression: "7 * 8 / 9",
			expected:   6.222222,
		},
		{
			expression: "(1 + 2) * 3",
			expected:   9,
		},
		{
			expression: "5 - (6 / 2) + 3",
			expected:   5,
		},
		{
			expression: "2 * (4 / ((1 + 2) + 1))",
			expected:   2,
		},
		{
			expression: "4 + 5.2",
			expected:   9.2,
		},
		{
			expression: "(3.2 + 3) / 2",
			expected:   3.1,
		},
		{
			expression: "-(-11-(1*20/2)-11/2*3)",
			expected:   37.5,
		},
		{
			expression:  "3.4 3 + 1",
			expectError: true,
		},
		{
			expression:  "5 - (6 / 0) + 3",
			expectError: true,
		},
		{
			expression:  "2 * (3 + 4",
			expectError: true,
		},
		{
			expression:  "2 * (3 + 4))",
			expectError: true,
		},
		{
			expression:  "",
			expectError: true,
		},
		{
			expression:  "4. 6 + 1",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			result, err := calculator.Eval(tt.expression)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
