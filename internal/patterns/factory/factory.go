package factory

type WorkCategory int

type IFactory interface {
	Create(creater string, title string, category WorkCategory) *Work
	createWork(creater string, title string, category WorkCategory) *Work
	registerWork(*Work)
}

type BaseFactory struct {
	IFactory
}

func (bf BaseFactory) Create(creater string, title string, category WorkCategory) *Work {
	w := bf.IFactory.createWork(creater, title, category)
	bf.IFactory.registerWork(w)

	return w
}
