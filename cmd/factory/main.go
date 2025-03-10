package main

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/creator"
)

func main() {
	bookCreator := creator.BookCreator{}
	movieCreator := creator.MovieCreator{}

	b1 := bookCreator.New("The Sirens of Titan", "Kurt Vonnegut")
	b2 := bookCreator.New("The Three-Body Problem", "Cixin Liu")
	m1 := movieCreator.New("About Time", "Richard Curtis")
	m2 := movieCreator.New("Forrest Gump", "Robert Lee Zemeckis")
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
