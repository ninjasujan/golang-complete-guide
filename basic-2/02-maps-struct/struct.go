package main

import (
	"fmt"
)

/**

	struct by deafult value type
**/

type Patient struct {
	patientList []string
}

type Doctor struct {
	patientList Patient
	number int
	doctorName string `required max:"100"`
	companion []string
}


func main() {
	doctor := Doctor {
		number: 12,
		doctorName: "Ryu",
		companion: []string {"Shaun", "Peiling", "Ninja"},
		patientList: Patient{
			patientList: []string{"1", "2", "3"},
		},
	}
	fmt.Println(doctor.companion)
}


