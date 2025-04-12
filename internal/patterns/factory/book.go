package factory

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/factory/work"
)

type BookFactory struct {
}

func (bf BookFactory) New(title, creator string) *work.Book {
	b := bf.create(title, creator)
	bf.register(b)

	return b
}

func (bf BookFactory) create(title, creator string) *work.Book {
	return &work.Book{
		Title:   title,
		Creator: creator,
	}
}

func (bf BookFactory) register(b *work.Book) *work.Book {
	fmt.Printf("registered book: %s\n", b.Title)

	return b
}
