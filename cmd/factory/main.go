package main

import (
	"fmt"
	"os"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

func main() {
	var rwf factory.RootWorkFactory = factory.RootWorkFactory{}
	// create factories
	bookFactory, err := rwf.Generate(work.BookCategory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	movieFactory, err := rwf.Generate(work.MovieCategory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b1 := bookFactory.Create("The Sirens of Titan", "Kurt Vonnegut")
	b2 := bookFactory.Create("The Three-Body Problem", "Cixin Liu")
	m1 := movieFactory.Create("About Time", "Richard Curtis")
	m2 := movieFactory.Create("Forrest Gump", "Robert Lee Zemeckis")
	// Output:
	// registered a Book: The Sirens of Titan
	// registered a Book: The Three-Body Problem
	// registered a Movie: About Time
	// registered a Movie: Forrest Gump

	b1.Display()
	b2.Display()
	m1.Display()
	m2.Display()
	// Output:
	// 【 Book 】
	//  Title: The Sirens of Titan
	//  Creater: Kurt Vonnegut
	// 【 Book 】
	//  Title: The Three-Body Problem
	//  Creater: Cixin Liu
	// 【 Movie 】
	//  Title: About Time
	//  Creater: Richard Curtis
	// 【 Movie 】
	//  Title: Forrest Gump
	//  Creater: Robert Lee Zemeckis
}
