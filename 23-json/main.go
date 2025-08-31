package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Timestamp struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	fmtTime := t.Time.Format(time.RFC3339)
	return json.Marshal(fmtTime)
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var fmtTime string
	if err := json.Unmarshal(data, &fmtTime); err != nil {
		return err
	}
	parsedTime, err := time.Parse(time.RFC3339, fmtTime)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

func main() {

	fmt.Println("JSON parsing example")

	p := Person{Name: "Alice", Age: 30}

	pJson, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println("JSON output:", string(pJson))

	var p2 Person
	err = json.Unmarshal(pJson, &p2)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Go struct output:", p2)

	t := Timestamp{Time: time.Now()}
	tJson, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Error marshalling Timestamp JSON:", err)
		return
	}

	fmt.Println("Timestamp JSON output:", string(tJson))

	var t2 Timestamp
	err = json.Unmarshal(tJson, &t2)
	if err != nil {
		fmt.Println("Error unmarshalling Timestamp JSON:", err)
		return
	}

	fmt.Println("Go struct output:", t2)
}
