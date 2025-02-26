package factory

type IWork interface {
	String() string
}

type Work struct {
	IWork
	Title string
}
