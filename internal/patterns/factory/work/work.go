package work

import "fmt"

type WorkCategory int

const (
	Undefined WorkCategory = iota
	BookCategory
	MovieCategory
)

func (wc WorkCategory) String() string {
	switch wc {
	case BookCategory:
		return "Book"
	case MovieCategory:
		return "Movie"
	default:
		return "Undefined"
	}
}

type IWork interface {
	Display()
}

type Work struct {
	Title    string
	Creater  string
	Category WorkCategory
	IWork
}

func (w Work) Display() {
	fmt.Printf("Category: %s, Title: %s, Creater: %s\n", w.Category.String(), w.Title, w.Creater)
}
