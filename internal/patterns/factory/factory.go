package factory

type IWorkFactory interface {
	Create(title, creater string) *Work
	createWork(title, creater string) *Work
	registerWork(w *Work)
}

type RootWorkFactory struct {
	Category WorkCategory
}

type WorkFactory struct {
	RootWorkFactory
	IWorkFactory
}

func (rwf RootWorkFactory) Create(wc WorkCategory) *WorkFactory {
	var factory IWorkFactory = rwf.RegisterFactory(wc)

	return &WorkFactory{
		RootWorkFactory: RootWorkFactory{Category: wc},
		IWorkFactory:    factory,
	}
}

func (rfc RootWorkFactory) RegisterFactory(wc WorkCategory) IWorkFactory {

}
