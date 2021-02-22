package usecase

import "fmt"

type Reception struct {
    next Department
}

func (r *Reception) Execute(p *Patient) {
    if p.RegistrationDone {
        fmt.Println("Patient registration already done")
        r.next.Execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.RegistrationDone = true
    r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
    r.next = next
}