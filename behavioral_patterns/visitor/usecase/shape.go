package usecase

type Shape interface {
    GetType() string
    Accept(Visitor)
}