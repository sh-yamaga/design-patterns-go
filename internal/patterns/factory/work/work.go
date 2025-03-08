package work

import "fmt"

type Category int

const (
	Undefined Category = iota
	Book
	Movie
)

func (c Category) String() string {
	switch c {
	case Book:
		return "Book"
	case Movie:
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
	Category Category
	IWork
}

func (w Work) Display() {
	fmt.Printf("【 %s 】\n Title: %s\n Creater: %s\n", w.Category.String(), w.Title, w.Creater)
}
