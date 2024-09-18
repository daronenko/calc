package main

import (
	"fmt"
	"os"

	"github.com/daronenko/calc/internal/calculator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You need to pass mathematical expression to calculate.\n")
		os.Exit(1)
	}
	expression := os.Args[1]

	result, err := calculator.Eval(expression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
