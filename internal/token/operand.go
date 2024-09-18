package token

import (
	"fmt"
	"strconv"
)

type Operand struct {
	value int
}

func NewOperand(value string) (*Operand, error) {
	if !isOperand(value) {
		return nil, fmt.Errorf("invalid operand: '%s'", value)
	}

	operand, _ := strconv.Atoi(value)
	return &Operand{operand}, nil
}

func (op *Operand) Value() int {
	return op.value
}

func isOperand(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}
