package category

type Work int

const (
	Undefined Work = iota
	Book
	Movie
)

func (w Work) String() string {
	switch w {
	case Book:
		return "Book"
	case Movie:
		return "Movie"
	default:
		return "Undefined"
	}
}
