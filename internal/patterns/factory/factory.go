package factory

import (
	"errors"
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type IWorkFactory interface {
	Create(title, creater string) *work.Work
	createWork(title, creater string) *work.Work
	registerWork(w *work.Work)
}

type RootWorkFactory struct {
	Category work.WorkCategory
}

type WorkFactory struct {
	RootWorkFactory
	IWorkFactory
}

func (rwf RootWorkFactory) Create(wc work.WorkCategory) *WorkFactory {
	factory, err := rwf.RegisterFactory(wc)
	if err != nil {
		fmt.Println(err)
	}

	return &WorkFactory{
		RootWorkFactory: RootWorkFactory{Category: wc},
		IWorkFactory:    factory,
	}
}

func (rfc RootWorkFactory) RegisterFactory(wc work.WorkCategory) (IWorkFactory, error) {
	switch wc {
	case work.BookCategory:
		return BookFactory{}, nil
	case work.MovieCategory:
		return MovieFactory{}, nil
	default:
		return nil, errors.New("unexpected WorkCategory was given")
	}
}
