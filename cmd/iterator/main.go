package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/iterator/fruit"
)

func main() {
	fb := fruit.NewBasket([]fruit.Fruit{})

	fb.Append(*fruit.NewFruit("Apple", "red"))
	fb.Append(*fruit.NewFruit("Orange", "orange"))
	fb.Append(*fruit.NewFruit("Grape", "purple"))
	fb.Append(*fruit.NewFruit("Banana", "yellow"))

	ittr := fb.Iterator()

	for ittr.HasNext() {
		f := ittr.Next()
		fmt.Printf("%s [%s]\n", f.Name(), f.Color())
	}

	// Output:
	// Apple [red]
	// Orange [orange]
	// Grape [purple]
	// Banana [yellow]
}
