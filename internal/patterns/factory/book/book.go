package book

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
)

type BookCategory factory.WorkCategory

const (
	Undefined BookCategory = iota
	Manga
	History
	SF
)

func (bc BookCategory) String() string {
	switch bc {
	case Undefined:
		return "undefined"
	case Manga:
		return "manga"
	case History:
		return "history"
	default:
		return "undefined"
	}
}

type Book struct {
	Author   string
	Title    string
	Category BookCategory
}

func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s, Category: %s", b.Title, b.Author, b.Category.String())
}
