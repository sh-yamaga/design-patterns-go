package factory

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/factory/work"
)

type MovieFactory struct {
}

func (mf MovieFactory) New(title, creator string) *work.Movie {
	m := mf.createMovie(title, creator)
	mf.registerMovie(m)

	return m
}

func (mf MovieFactory) createMovie(title, creator string) *work.Movie {
	return &work.Movie{
		Title:   title,
		Creator: creator,
	}
}

func (mf MovieFactory) registerMovie(m *work.Movie) *work.Movie {
	fmt.Printf("registered movie: %s\n", m.Title)

	return m
}
