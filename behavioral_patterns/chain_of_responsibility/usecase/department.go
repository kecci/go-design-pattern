package usecase

type Department interface {
    Execute(*Patient)
    SetNext(Department)
}