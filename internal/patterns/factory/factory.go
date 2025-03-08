package factory

import (
	"errors"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"
)

type IWorkFactory interface {
	Create(title, creater string) *work.Work
	createWork(title, creater string) *work.Work
	registerWork(w *work.Work)
}

type RootWorkFactory struct {
	Category work.Category
}

func (rwf RootWorkFactory) Generate(wc work.Category) (IWorkFactory, error) {
	factory, err := rwf.registerFactory(wc)
	if err != nil {
		return nil, err
	}

	return factory, nil
}

func (rfc RootWorkFactory) registerFactory(wc work.Category) (IWorkFactory, error) {
	switch wc {
	case work.Book:
		bf := BookFactory{}
		// register method to interface
		bf.Factory.IWorkFactory = bf

		return bf, nil
	case work.Movie:
		mf := MovieFactory{}
		// register method to interface
		mf.Factory.IWorkFactory = mf

		return mf, nil
	default:
		return nil, errors.New("unexpected WorkCategory was given")
	}
}

type Factory struct {
	IWorkFactory
}

func (f Factory) Create(title, creater string) *work.Work {
	var w *work.Work = f.IWorkFactory.createWork(title, creater)
	f.IWorkFactory.registerWork(w)

	return w
}
