package usecase

type Director struct {
    builder iBuilder
}

func NewDirector(b iBuilder) *Director {
    return &Director{
        builder: b,
    }
}

func (d *Director) SetBuilder(b iBuilder) {
    d.builder = b
}

func (d *Director) BuildHouse() House {
    d.builder.setDoorType()
    d.builder.setWindowType()
    d.builder.setNumFloor()
    return d.builder.getHouse()
}