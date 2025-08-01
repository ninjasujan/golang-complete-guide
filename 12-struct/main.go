// This file demonstrates the use of structs in Go, including creating instances and updating fields via a pointer.

package main

import (
	"fmt"
)

func main() {

	fmt.Println("[Go User Defined Type Basics]")

	var person = Person{Name: "John", JobType: "Engineer", Salary: 15000}

	person2 := Person{Name: "Jane", JobType: "Doctor", Salary: 20000}

	fmt.Println(person, person2)

	UpdateSalary(&person)

	fmt.Println(person)

	person.Display()

	employee := Employee{"Sujan", 20000}

	fmt.Println("[Employee Name]", employee.string)

	fmt.Println("[Employee Salary]", employee.int)

	student := createStudent("John", "20th", 3.5, "Computer Science")

	fmt.Println(student)
	fmt.Println(student.Person)

}

func createStudent(name string, class string, gpa float32, course string) Student {
	return Student{
		Person: &Person{
			Name: name,
		},
		course: course,
		gpa:    gpa,
		class:  class,
	}
}

type Person struct {
	Name    string
	JobType string
	Salary  uint
}

type Employee struct {
	string
	int
}

type Student struct {
	class  string
	course string
	gpa    float32
	*Person
}

func UpdateSalary(person *Person) {
	person.Salary = 10000
}

func (person *Person) Display() {
	fmt.Printf("Name: %s, JobType: %s, Salary: %d\n", person.Name, person.JobType, person.Salary)
}
