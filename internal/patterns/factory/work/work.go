package work

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/category"
)

type IWork interface {
	Display()
}

type Work struct {
	Title    string
	Creater  string
	Category category.Work
	IWork
}

func (w Work) Display() {
	fmt.Printf("【 %s 】\n Title: %s\n Creater: %s\n", w.Category.String(), w.Title, w.Creater)
}
