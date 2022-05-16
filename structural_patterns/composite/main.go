package main

import "github.com/kecci/go-design-pattern/structural_patterns/composite/usecase"

func main() {
	file1 := &usecase.File{Name: "File1"}
	file2 := &usecase.File{Name: "File2"}
	file3 := &usecase.File{Name: "File3"}

	folder1 := &usecase.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &usecase.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
