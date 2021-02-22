package usecase

import "fmt"

type Doctor struct {
    next Department
}

func (d *Doctor) Execute(p *Patient) {
    if p.DoctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.Execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.DoctorCheckUpDone = true
    d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
    d.next = next
}