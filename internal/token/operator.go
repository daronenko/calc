package token

import (
	"fmt"
)

type Operation func(lhs, rhs *Operand) (*Operand, error)

var operators = map[string]*Operator{
	"+": {
		value:      "+",
		precedence: 10,
		operation: func(lhs, rhs *Operand) (*Operand, error) {
			result, _ := NewOperand(fmt.Sprintf("%f", lhs.Value()+rhs.Value()))
			return result, nil
		},
	},
	"-": {
		value:      "-",
		precedence: 10,
		operation: func(lhs, rhs *Operand) (*Operand, error) {
			result, _ := NewOperand(fmt.Sprintf("%f", lhs.Value()-rhs.Value()))
			return result, nil
		},
	},
	"*": {
		value:      "*",
		precedence: 20,
		operation: func(lhs, rhs *Operand) (*Operand, error) {
			result, _ := NewOperand(fmt.Sprintf("%f", lhs.Value()*rhs.Value()))
			return result, nil
		},
	},
	"/": {
		value:      "/",
		precedence: 20,
		operation: func(lhs, rhs *Operand) (*Operand, error) {
			if rhs.Value() == 0 {
				return nil, fmt.Errorf("division by zero")
			}

			result, _ := NewOperand(fmt.Sprintf("%f", lhs.Value()/rhs.Value()))
			return result, nil
		},
	},
}

type Operator struct {
	value      string
	precedence int
	operation  Operation
}

func NewOperator(value string) (*Operator, error) {
	op, exists := operators[value]
	if !exists {
		return nil, fmt.Errorf("invalid operator: `%s`", value)
	}
	return op, nil
}

func (op *Operator) Precedence() int {
	return op.precedence
}

func (op *Operator) Call(a, b *Operand) (*Operand, error) {
	result, err := op.operation(a, b)
	if err != nil {
		return nil, fmt.Errorf("operation cannot be performed: %v", err)
	}

	return result, nil
}

func isOperator(value string) bool {
	_, exists := operators[value]
	return exists
}
