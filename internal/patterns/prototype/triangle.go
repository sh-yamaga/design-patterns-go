package prototype

import "fmt"

// Triangle represents an equilateral triangle with a given height.
type Triangle struct {
	Height int
}

func (t *Triangle) Draw() {
	slice := make([]struct{}, t.Height)
	for i := range slice {
		// Print leading spaces
		spaceSlice := make([]struct{}, t.Height-i-1)
		for range spaceSlice {
			fmt.Print(" ")
		}
		// Print stars for the triangle
		starSlice := make([]struct{}, 2*i+1)
		for range starSlice {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (t *Triangle) Clone() Shape {
	return &Triangle{Height: t.Height}
}
