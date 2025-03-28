package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/visitor"
)

func main() {
	// Add 2 shapes to the Element slice
	elements := []visitor.Element{
		&visitor.Circle{Radius: 3.0},
		&visitor.Rectangle{Width: 4.0, Height: 5.5},
	}

	// Create 2 visitors
	areaVisitor := &visitor.AreaVisitor{}
	displayVisitor := &visitor.DisplayVisitor{}

	fmt.Println("=== Visitor Displaying Element Information ===")
	for _, e := range elements {
		e.Accept(displayVisitor)
	}

	fmt.Println()

	fmt.Println("=== Visitor Calculating Element Areas ===")
	for _, e := range elements {
		e.Accept(areaVisitor)
	}
}
