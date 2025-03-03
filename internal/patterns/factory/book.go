package factory

import "fmt"

type BookFactory struct {
	WorkFactory
}

func (bf BookFactory) Create(title, creater string) *Work {
	var book *Work = bf.CreateWork(title, creater)
	bf.RegisterWork(book)

	return book
}

func (bf BookFactory) CreateWork(title, creater string) *Work {
	return &Work{
		Title:    title,
		Creater:  creater,
		Category: bf.WorkFactory.Category,
	}
}

func (bf BookFactory) RegisterWork(w *Work) {
	fmt.Println("registered book: " + w.Title)
}
