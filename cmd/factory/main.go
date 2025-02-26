package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/book"
)

func main() {
	bf := book.NewBookFactory()

	b1 := bf.Create("Kurt Vonnegut", "The Sirens of Titan", factory.WorkCategory(book.SF))
	b2 := bf.Create("Osamu Tezuka", "Phoenix", factory.WorkCategory(book.Manga))
	b3 := bf.Create("Yuval Noah Harari", "Sapiens: A Brief History of Humankind", factory.WorkCategory(book.History))

	fmt.Println(b1.String())
	fmt.Println(b2.String())
	fmt.Println(b3.String())
}
