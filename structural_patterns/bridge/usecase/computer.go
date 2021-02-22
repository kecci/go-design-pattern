package usecase

type Computer interface {
    Print()
    SetPrinter(Printer)
}