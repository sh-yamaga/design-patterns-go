package factory

import (
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/book"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/movie"
)

type IWorkFactory interface {
	BookFactory() book.BookFactory
	MovieFactory() movie.MovieFactory
}

type WorkFactory struct {
	IWorkFactory
}

func (wf WorkFactory) BookFactory() book.BookFactory {
	return book.NewBookFacotry()
}

func (wf WorkFactory) MovieFactory() movie.MovieFactory {
	return movie.NewMovieFacotry()
}
