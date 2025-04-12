package factory

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/factory/work"
)

type MovieFactory struct {
}

func (mf MovieFactory) New(title, creator string) *work.Movie {
	m := mf.create(title, creator)
	mf.register(m)

	return m
}

func (mf MovieFactory) create(title, creator string) *work.Movie {
	return &work.Movie{
		Title:   title,
		Creator: creator,
	}
}

func (mf MovieFactory) register(m *work.Movie) *work.Movie {
	fmt.Printf("registered movie: %s\n", m.Title)

	return m
}
