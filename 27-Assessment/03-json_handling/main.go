package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func main() {

	user := &User{Name: "Sujan", Email: "sujan@gmail.com", Address: "2012, Singasandra Bengaluru, karnataka"}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println("[error marshaling json]")
	}

	fmt.Println("[JSON Marshaling]", string(jsonData))

	var user2 User
	err = json.Unmarshal(jsonData, &user2)

	if err != nil {
		fmt.Println("[error unmarshaling json]")
	}

	fmt.Println("[JSON Unmarshaling]", user2)

}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		Email string `json:"email"`
	}{
		Alias: (*Alias)(u),
		Email: "[REDACTED]",
	})
}
