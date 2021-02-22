package main

import (
	"fmt"

	"github.com/kecci/design-pattern-go/creational_patterns/prototype/usecase"
)

func main() {
	file1 := &usecase.File{Name: "File1"}
	file2 := &usecase.File{Name: "File2"}
	file3 := &usecase.File{Name: "File3"}

	folder1 := &usecase.Folder{
		Children: []usecase.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &usecase.Folder{
		Children: []usecase.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
