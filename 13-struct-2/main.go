package main

import "fmt"

func main() {
	person := createPerson("Sujan", "sujan@gmail.com", struct {
		City    string
		PinCode uint16
		Country string
	}{
		City:    "Kathmandu",
		PinCode: 44600,
		Country: "Nepal",
	}, struct {
		Facebook  string
		Instagram string
		Twitter   string
	}{
		Facebook:  "facebook",
		Instagram: "instagram",
		Twitter:   "twitter",
	})

	// Now you can access the fields
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("City: %s\n", person.Address.City)
	fmt.Printf("Instagram: %s\n", person.SocialMedia.Instagram)

}

func createPerson(name string, email string, address struct {
	City    string
	PinCode uint16
	Country string
}, socialMedia struct {
	Facebook  string
	Instagram string
	Twitter   string
}) *Person {
	return &Person{
		Name:        name,
		Email:       email,
		Address:     address,
		SocialMedia: socialMedia,
	}
}

type Person struct {
	Name    string
	Email   string
	Address struct {
		City    string
		PinCode uint16
		Country string
	}
	SocialMedia struct {
		Facebook  string
		Instagram string
		Twitter   string
	}
}
