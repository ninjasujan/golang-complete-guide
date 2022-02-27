package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

// receiver function - struct
func (p person) print() {
	fmt.Printf("%+v", p)
}


// update struct member
func (p *person) updatePerson(email string) {
	(*p).contact.email = email
};

func main() {
	alex := person{
		firstName: "Alex", 
		lastName: "Anderson",
		contact: contactInfo {
			zipCode: 576234,
			email: "alex@gmail.com",
		},
	}
	alex.print()
	// alexPointer := &alex;
	alex.updatePerson("alex@appyhigh.com") // auto conversion of type pointer
	alex.print()
}