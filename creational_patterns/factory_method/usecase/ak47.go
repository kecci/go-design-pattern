package usecase

type Ak47 struct {
    Gun
}

func NewAk47() IGun {
    return &Ak47{
        Gun: Gun{
            Name:  "AK47 gun",
            Power: 4,
        },
    }
}