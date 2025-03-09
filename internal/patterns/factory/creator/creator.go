package creator

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"

type IWorkCreator interface {
	Create(title, creator string) work.IWork
	createWork(title, creator string) *work.Work
	registerWork(w *work.Work) work.IWork
}

type BaseCreator struct {
	IWorkCreator
}

func (bc BaseCreator) Create(title, creator string) work.IWork {
	var w *work.Work = bc.IWorkCreator.createWork(title, creator)
	return bc.IWorkCreator.registerWork(w)
}
