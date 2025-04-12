package main

import "github.com/yamaga-shu/design-patterns-go/internal/patterns/factory"

func main() {
	bookFactory := &factory.BookFactory{}
	movieFactory := &factory.MovieFactory{}

	b1 := bookFactory.New("The Sirens of Titan", "Kurt Vonnegut")
	b2 := bookFactory.New("The Three-Body Problem", "Cixin Liu")
	m1 := movieFactory.New("About Time", "Richard Curtis")
	m2 := movieFactory.New("Forrest Gump", "Robert Lee Zemeckis")
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
	// 【Book】
	// Title: The Sirens of Titan
	// Creator: Kurt Vonnegut
	// 【Book】
	// Title: The Three-Body Problem
	// Creator: Cixin Liu
	// 【Movie】
	// Title: About Time
	// Creator: Richard Curtis
	// 【Movie】
	// Title: Forrest Gump
	// Creator: Robert Lee Zemeckis
}
