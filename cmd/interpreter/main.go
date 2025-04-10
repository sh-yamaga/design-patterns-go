package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/interpreter"
)

func main() {
	// Generate Expression
	/// Expression which checks if a "milk" is contained.
	milk := &interpreter.TerminalExpression{
		Word: "milk",
	}

	/// Expression which checks if a "sugar" is contained.
	sugar := &interpreter.TerminalExpression{
		Word: "sugar",
	}

	//// Expression which checks if either "milk" or "sugar" is contained
	milkOrSugar := &interpreter.OrExpression{
		Expr1: milk,
		Expr2: sugar,
	}

	//// Expression which checks if "milk" and "sugar" are both contained
	milkAndSugar := &interpreter.AndExpression{
		Expr1: milk,
		Expr2: sugar,
	}

	// testing
	texts := []string{
		"Coffee with milk",
		"Coffee with sugar",
		"Coffee with milk and sugar",
		"Coffee only",
	}

	for _, text := range texts {
		fmt.Println("===", text, "===")
		fmt.Println("Contains \"milk\" or \"sugar\"?:", milkOrSugar.Interpret(text))
		fmt.Println("Contains \"milk\" and \"sugar\"?:", milkAndSugar.Interpret(text))
	}
	// Output:
	// === Coffee with milk ===
	// Contains "milk" or "sugar"?: true
	// Contains "milk" and "sugar"?: false
	// === Coffee with sugar ===
	// Contains "milk" or "sugar"?: true
	// Contains "milk" and "sugar"?: false
	// === Coffee with milk and sugar ===
	// Contains "milk" or "sugar"?: true
	// Contains "milk" and "sugar"?: true
	// === Coffee only ===
	// Contains "milk" or "sugar"?: false
	// Contains "milk" and "sugar"?: false
}
