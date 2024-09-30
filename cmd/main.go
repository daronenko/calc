package main

import (
	"log"
	"os"

	"github.com/daronenko/calc/internal/calculator"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	if len(os.Args) < 2 {
		log.Fatalf("You need to pass mathematical expression to calculate.\n")
	}
	expression := os.Args[1]

	result, err := calculator.Eval(expression)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	log.Println(result)
}
