package composite

import (
	"fmt"
	"strings"
)

type File struct {
	Name  string
	Level int // 階層レベル
}

func (f *File) Show() {
	indent := strings.Repeat("  ", f.Level)
	fmt.Println(indent + f.Name)
}
