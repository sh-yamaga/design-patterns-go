package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/singleton"
)

func main() {
	s1 := singleton.New()
	s2 := singleton.New()

	// This is compile error
	// s1.data = "another value"

	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}
	// Output:
	// s1 == s2
}
