package composite

import (
	"fmt"
	"strings"
)

type File struct {
	Name  string
	Level int
}

func (f *File) Show() {
	indent := strings.Repeat("  ", f.Level)
	fmt.Println(indent + f.Name)
}
