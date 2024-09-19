package token

import (
	"fmt"
	"strconv"
)

type Operand struct {
	value float64
}

func NewOperand(value string) (*Operand, error) {
	if !isOperand(value) {
		return nil, fmt.Errorf("invalid operand: '%s'", value)
	}

	operand, _ := strconv.ParseFloat(value, 64)
	return &Operand{operand}, nil
}

func (op *Operand) Value() float64 {
	return op.value
}

func isOperand(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}
