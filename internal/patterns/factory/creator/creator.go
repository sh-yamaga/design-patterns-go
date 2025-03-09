package creator

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/work"

type IWorkCreator interface {
	Create(title, creater string) *work.Work
	createWork(title, creater string) *work.Work
	registerWork(w *work.Work)
}

type BaseCreator struct {
	IWorkCreator
}

func (bc BaseCreator) Create(title, creater string) *work.Work {
	var w *work.Work = bc.IWorkCreator.createWork(title, creater)
	bc.IWorkCreator.registerWork(w)

	return w
}
