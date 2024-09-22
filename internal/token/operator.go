package token

import (
	"fmt"
)

type Operation func(...Operand) (*Operand, error)

var operators = map[string]*Operator{
	"+": {
		value:      "+",
		precedence: 10,
		operation: func(args ...Operand) (*Operand, error) {
			if len(args) > 2 {
				return nil, fmt.Errorf("invalid count of arguments for operation '+'")
			}

			var result *Operand

			if len(args) == 1 {
				result, _ = NewOperand(toString(args[0].Value()))
			} else {
				result, _ = NewOperand(toString(args[0].Value() + args[1].Value()))
			}

			return result, nil
		},
	},
	"-": {
		value:      "-",
		precedence: 10,
		operation: func(args ...Operand) (*Operand, error) {
			if len(args) > 2 {
				return nil, fmt.Errorf("invalid count of arguments for operation '-'")
			}

			var result *Operand

			if len(args) == 1 {
				result, _ = NewOperand(toString(-args[0].Value()))
			} else {
				result, _ = NewOperand(toString(args[0].Value() - args[1].Value()))
			}

			return result, nil
		},
	},
	"*": {
		value:      "*",
		precedence: 20,
		operation: func(args ...Operand) (*Operand, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("'*' operator takes two arguments")
			}

			result, _ := NewOperand(toString(args[0].Value() * args[1].Value()))
			return result, nil
		},
	},
	"/": {
		value:      "/",
		precedence: 20,
		operation: func(args ...Operand) (*Operand, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("'/' operator takes two arguments")
			}

			if args[1].Value() == 0 {
				return nil, fmt.Errorf("division by zero")
			}

			result, _ := NewOperand(toString(args[0].Value() / args[1].Value()))
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

func (op *Operator) Call(args ...Operand) (*Operand, error) {
	result, err := op.operation(args...)
	if err != nil {
		return nil, fmt.Errorf("operation cannot be performed: %v", err)
	}

	return result, nil
}

func (op *Operator) Operation() string {
	return op.value
}

func isOperator(value string) bool {
	_, exists := operators[value]
	return exists
}

func toString(value float64) string {
	return fmt.Sprintf("%f", value)
}
