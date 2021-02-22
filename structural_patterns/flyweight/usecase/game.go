package usecase

type Game struct {
}

func NewGame() Game {
	return Game{}
}

func (g *Game) AddTerrorist(dressType string) {
	DressFactorySingleInstance.DressMap[dressType] = NewTerroristDress()
}

func (g *Game) AddCounterTerrorist(dressType string) {
	DressFactorySingleInstance.DressMap[dressType] = NewCounterTerroristDress()
}
