package usecase

type Originator struct {
    State string
}

func (e *Originator) CreateMemento() *Memento {
    return &Memento{state: e.State}
}

func (e *Originator) Restorememento(m *Memento) {
    e.State = m.GetSavedState()
}

func (e *Originator) SetState(state string) {
    e.State = state
}

func (e *Originator) GetState() string {
    return e.State
}