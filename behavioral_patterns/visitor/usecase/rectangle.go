package usecase

type Rectangle struct {
    L int
    B int
}

func (t *Rectangle) Accept(v Visitor) {
    v.VisitForrectangle(t)
}

func (t *Rectangle) GetType() string {
    return "rectangle"
}