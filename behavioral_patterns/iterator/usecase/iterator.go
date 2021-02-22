package usecase

type Iterator interface {
    HasNext() bool
    GetNext() *User
}