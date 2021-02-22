package usecase

type Mediator interface {
    CanArrive(Train) bool
    NotifyAboutDeparture()
}