package main

import "fmt"

func main() {
	// create original streem
	originalStream := &decorator.Stream{Data: "hello, decorator pattern"}

	//  add filter streem to uppercase
	upperStream := &decorator.UpperCaseFilter{
		Wrapped: originalStream,
	}

	// before filtering
	fmt.Println("original Stream: ", originalStream.Read())
	// Output:
	// hello, decorator pattern

	// after filtering
	fmt.Println("Filtered Stream: ", upperStream.Read())
	// Output:
	// HELLO, DECORATOR PATTERN
}
