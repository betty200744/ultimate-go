package chain_of_responsibility

import (
	"fmt"
	"testing"
)

func Test_patient(t *testing.T) {
	reception := &reception{}
	doctor := &doctor{}
	medical := &medical{}
	cashier := &cashier{}
	patient := &patient{
		name:              "nancy",
		registrationDone:  false,
		doctorCheckUpDone: false,
		medicineDone:      false,
		paymentDone:       false,
	}
	reception.setNext(doctor)
	doctor.setNext(medical)
	medical.setNext(cashier)
	reception.execute(patient)
	fmt.Println(patient)
}
