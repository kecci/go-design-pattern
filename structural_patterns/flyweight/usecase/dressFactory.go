package usecase

import "fmt"

const (
    //TerroristDressType terrorist dress type
    TerroristDressType = "tDress"
    //CounterTerrroristDressType terrorist dress type
    CounterTerrroristDressType = "ctDress"
)

var (
    DressFactorySingleInstance = &DressFactory{
        DressMap: make(map[string]Dress),
    }
)

type DressFactory struct {
    DressMap map[string]Dress
}

func (d *DressFactory) GetDressByType(dressType string) (Dress, error) {
    if d.DressMap[dressType] != nil {
        return d.DressMap[dressType], nil
    }

    if dressType == TerroristDressType {
        d.DressMap[dressType] = NewTerroristDress()
        return d.DressMap[dressType], nil
    }
    if dressType == CounterTerrroristDressType {
        d.DressMap[dressType] = NewCounterTerroristDress()
        return d.DressMap[dressType], nil
    }

    return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
    return DressFactorySingleInstance
}