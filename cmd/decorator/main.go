package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/decorator"
)

func main() {
	// create original streem
	originalStream := &decorator.Stream{Data: "hello, decorator pattern"}

	//  add filter streem to uppercase
	upperStream := &decorator.UpperCaseFilter{
		Wrapped: originalStream,
	}

	// before filtering
	fmt.Println("Original Stream:", originalStream.Read())
	// Output:
	// Original Stream: hello, decorator pattern

	// after filtering
	fmt.Println("Filtered Stream:", upperStream.Read())
	// Output:
	// Filtered Stream: HELLO, DECORATOR PATTERN
}
