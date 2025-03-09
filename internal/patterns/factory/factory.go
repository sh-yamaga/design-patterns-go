package factory

import (
	"errors"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/creator"
)

type CreatorFactory struct {
	Category category.Work
}

func (cf CreatorFactory) NewCreator(wc category.Work) (creator.IWorkCreator, error) {
	factory, err := cf.registerCreator(wc)
	if err != nil {
		return nil, err
	}

	return factory, nil
}

func (cf CreatorFactory) registerCreator(wc category.Work) (creator.IWorkCreator, error) {
	switch wc {
	case category.Book:
		bc := creator.BookCreator{}
		bc.BaseCreator.IWorkCreator = bc

		return bc, nil
	case category.Movie:
		mc := creator.MovieCreator{}
		mc.BaseCreator.IWorkCreator = mc

		return mc, nil
	default:
		return nil, errors.New("unexpected category.Work was given")
	}
}
