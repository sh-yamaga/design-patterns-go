package main

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/composite"

func main() {
	// create root dir
	root := &composite.Directory{Name: "root", Level: 0}

	// create child dir
	dir2 := &composite.Directory{Name: "dir2"}
	dir1 := &composite.Directory{Name: "dir1"}

	// create files
	file1 := &composite.File{Name: "file1.txt"}
	file2 := &composite.File{Name: "file2.txt"}
	file3 := &composite.File{Name: "file3.txt"}

	// construct
	// root -> dir2 -> (dir1 -> (file1, file2), file3)

	root.AddDir(dir2)

	dir2.AddDir(dir1)

	dir1.AddFile(file1)
	dir1.AddFile(file2)

	dir2.AddFile(file3)

	// show tree
	root.Show()
	// Output:
	// root
	//   dir2
	//     dir1
	//       file1.txt
	//       file2.txt
	//     file3.txt
}
