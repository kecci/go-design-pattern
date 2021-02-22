package usecase

import "fmt"

type Item struct {
    observerList []Observer
    name         string
    inStock      bool
}

func NewItem(name string) *Item {
    return &Item{
        name: name,
    }
}
func (i *Item) UpdateAvailability() {
    fmt.Printf("Item %s is now in stock\n", i.name)
    i.inStock = true
    i.NotifyAll()
}
func (i *Item) Register(o Observer) {
    i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o Observer) {
    i.observerList = RemoveFromslice(i.observerList, o)
}

func (i *Item) NotifyAll() {
    for _, observer := range i.observerList {
        observer.Update(i.name)
    }
}

func RemoveFromslice(observerList []Observer, observerToRemove Observer) []Observer {
    observerListLength := len(observerList)
    for i, observer := range observerList {
        if observerToRemove.GetID() == observer.GetID() {
            observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
            return observerList[:observerListLength-1]
        }
    }
    return observerList
}