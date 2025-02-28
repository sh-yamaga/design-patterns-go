package factory

type WorkCategory int

const (
	Undefined WorkCategory = iota
	Book
	Movie
)

type IWork interface {
	Display()
}

type Work struct {
	Title    string
	Creater  string
	Category WorkCategory
	IWork
}
