package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/visitor"
	"github.com/yamaga-shu/design-patterns-go/internal/patterns/visitor/element"
)

func main() {
	// Create 2 shapes to the Element slice
	elements := []element.Element{
		&element.Circle{Radius: 3.0},
		&element.Rectangle{Width: 4.0, Height: 5.5},
	}

	// Create visitor which display each information
	displayVisitor := &visitor.DisplayVisitor{}

	fmt.Println("=== Visitor Displaying Element Information ===")
	for _, e := range elements {
		e.Accept(displayVisitor)
	}
	// Output:
	// === Visitor Displaying Element Information ===
	// DisplayVisitor: Circle (Radius=3.00)
	// DisplayVisitor: Rectangle (Width=4.00, Height=5.50)

	fmt.Println()

	// Create visitor which calculates the area of each shape
	areaVisitor := &visitor.AreaVisitor{}

	fmt.Println("=== Visitor Calculating Element Areas ===")
	for _, e := range elements {
		e.Accept(areaVisitor)
	}
	// Output:
	// === Visitor Calculating Element Areas ===
	// AreaVisitor: Area of Circle = 28.27
	// AreaVisitor: Area of Rectangle = 22.00
}
