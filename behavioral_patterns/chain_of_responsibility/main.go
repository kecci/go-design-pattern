package main

import (
	"fmt"

	"github.com/kecci/design-pattern-go/behavioral_patterns/chain_of_responsibility/usecase"
)

func main() {

	cashier := &usecase.Cashier{}

	//Set next for medical department
	medical := &usecase.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &usecase.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &usecase.Reception{}
	reception.SetNext(doctor)

	patient := &usecase.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)

	fmt.Printf("RegistrationDone %v \n", patient.RegistrationDone)
	fmt.Printf("DoctorCheckUpDone %v \n", patient.DoctorCheckUpDone)
	fmt.Printf("MedicineDone %v \n", patient.MedicineDone)
	fmt.Printf("PaymentDone %v \n", patient.PaymentDone)
}
