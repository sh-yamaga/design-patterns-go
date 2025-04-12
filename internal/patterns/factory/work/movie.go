package work

import "fmt"

type Movie struct {
	Title   string
	Creator string
}

func (m Movie) Display() {
	fmt.Println("【Movie】")
	fmt.Println("Title:", m.Title)
	fmt.Println("Creator:", m.Creator)
}
