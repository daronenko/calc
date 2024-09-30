package token

import "fmt"

type Bracket struct {
	value string
}

func NewBracket(value string) (*Bracket, error) {
	if !isBracket(value) {
		return nil, fmt.Errorf("invalid bracket: '%s'", value)
	}

	return &Bracket{value}, nil
}

func (b *Bracket) IsOpeningBracket() bool {
	return b.value == "("
}

func (b *Bracket) IsClosingBracket() bool {
	return b.value == ")"
}

func isBracket(char string) bool {
	return char == "(" || char == ")"
}
