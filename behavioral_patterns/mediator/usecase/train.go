package usecase

type Train interface {
    Arrive()
    Depart()
    PermitArrival()
}