package usecase

type Caretaker struct {
    MementoArray []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
    c.MementoArray = append(c.MementoArray, m)
}

func (c *Caretaker) GetMenento(index int) *Memento {
    return c.MementoArray[index]
}