package usecase

type Observer interface {
    Update(string)
    GetID() string
}