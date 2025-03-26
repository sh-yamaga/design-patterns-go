package composite

import (
	"fmt"
	"strings"
)

type Directory struct {
	Name     string
	Level    int
	children []FileSystemComponent
}

func (d *Directory) Show() {
	indent := strings.Repeat("  ", d.Level)
	fmt.Println(indent + d.Name)
	for _, child := range d.children {
		child.Show()
	}
}

func (d *Directory) setLevelRecursive(level int) {
	d.Level = level
	for _, child := range d.children {
		switch c := child.(type) {
		case *File:
			c.Level = level + 1
		case *Directory:
			c.setLevelRecursive(level + 1)
		}
	}
}

func (d *Directory) AddFile(file *File) {
	file.Level = d.Level + 1
	d.children = append(d.children, file)
}

func (d *Directory) AddDir(dir *Directory) {
	dir.setLevelRecursive(d.Level + 1)
	d.children = append(d.children, dir)
}
