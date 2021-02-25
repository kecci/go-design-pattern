package usecase

type Visitor interface {
    VisitForSquare(*Square)
    VisitForCircle(*Circle)
    VisitForrectangle(*Rectangle)
}