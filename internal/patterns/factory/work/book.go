package work

import "fmt"

type Book struct {
	Title   string
	Creator string
}

func (b Book) Display() {
	fmt.Println("【Book】")
	fmt.Println("Title:", b.Title)
	fmt.Println("Creator:", b.Creator)
}
