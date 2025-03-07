package main

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

func main() {
	var rwf factory.RootWorkFactory = factory.RootWorkFactory{}
	// create factories
	var bookFactory *factory.WorkFactory = rwf.Create(work.BookCategory)
	var movieFactory *factory.WorkFactory = rwf.Create(work.MovieCategory)

	// create books
	b1 := bookFactory.IWorkFactory.Create("The Sirens of Titan", "Kurt Vonnegut")
	b2 := bookFactory.IWorkFactory.Create("The Three-Body Problem", "Cixin Liu")
	// Output:
	// registered a Book: The Sirens of Titan
	// registered a Book: The Three-Body Problem

	// create movies
	m1 := movieFactory.IWorkFactory.Create("About Time", "Richard Curtis")
	m2 := movieFactory.IWorkFactory.Create("Forrest Gump", "Robert Lee Zemeckis")
	// Output:
	// registered a Movie: About Time
	// registered a Movie: Forrest Gump

	// display books
	b1.Display()
	b2.Display()
	// Output:
	// Category: Book, Title: The Sirens of Titan, Creater: Kurt Vonnegut
	// Category: Book, Title: The Three-Body Problem, Creater: Cixin Liu

	// display movies
	m1.Display()
	m2.Display()
	// Output:
	// Category: Movie, Title: About Time, Creater: Richard Curtis
	// Category: Movie, Title: Forrest Gump, Creater: Robert Lee Zemeckis
}
