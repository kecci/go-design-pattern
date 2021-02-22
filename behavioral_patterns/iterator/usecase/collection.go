package usecase

type Collection interface {
    CreateIterator() Iterator
}
