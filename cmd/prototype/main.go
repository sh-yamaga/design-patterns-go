package main

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/prototype"
)

func main() {
	triangle := &prototype.Triangle{Height: 4}
	square := &prototype.Square{Side: 5}

	triangle.Draw()
	// Output:
	//    *
	//   ***
	//  *****
	// *******
	square.Draw()
	// Output:
	// *****
	// *   *
	// *   *
	// *   *
	// *****

	clonedTriangle := triangle.Clone()
	clonedSquare := square.Clone()

	clonedTriangle.Draw()
	// Output:
	//    *
	//   ***
	//  *****
	// *******
	clonedSquare.Draw()
	// Output:
	// *****
	// *   *
	// *   *
	// *   *
	// *****
}
