package main

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/flyweight"
)

func main() {
	tf := flyweight.NewTreeFactory()

	trees := make([]*flyweight.Tree, 0)

	// create 100 hardwood trees
	for i := 0; i < 100; i++ {
		tt := tf.GetTreeType("Hardwood", "LightGreen", "BordLeaf")
		trees = append(trees, &flyweight.Tree{
			X:    i * 3,
			Y:    i * 2,
			Type: tt,
		})
	}

	// create 100 softwood trees
	for i := 0; i < 100; i++ {
		tt := tf.GetTreeType("Hardwood", "DeepGreen", "NeedleLeaf")
		trees = append(trees, &flyweight.Tree{
			X:    i * 2,
			Y:    i * 3,
			Type: tt,
		})
	}

	// draw
	for _, tree := range trees {
		tree.Draw()
	}
	// Output:
	// Name: Hardwood, Color: LightGreen, Texture: BordLeaf
	// Coordinate: (0, 0)
	// Name: Hardwood, Color: LightGreen, Texture: BordLeaf
	// Coordinate: (3, 2)
	// ...
	// ...
	// Name: Hardwood, Color: DeepGreen, Texture: NeedleLeaf
	// Coordinate: (0, 0)
	// Name: Hardwood, Color: DeepGreen, Texture: NeedleLeaf
	// Coordinate: (2, 3)
	// ...
	// ...
}
