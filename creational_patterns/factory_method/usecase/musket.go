package usecase

type Musket struct {
    Gun
}

func NewMusket() IGun {
    return &Musket{
        Gun: Gun{
            Name:  "Musket gun",
            Power: 1,
        },
    }
}