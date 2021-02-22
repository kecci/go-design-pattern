package usecase

import "fmt"

type Cashier struct {
    next Department
}

func (c *Cashier) Execute(p *Patient) {
    if p.PaymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
	p.PaymentDone = true
}

func (c *Cashier) SetNext(next Department) {
    c.next = next
}