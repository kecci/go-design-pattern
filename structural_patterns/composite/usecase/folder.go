package usecase

import "fmt"

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c)
}
