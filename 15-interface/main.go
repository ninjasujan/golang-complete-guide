package main

import "fmt"

func main() {

	employee := Employee{
		name:   "John",
		salary: 10000,
		bonus:  2000,
	}
	pay := employee.CalPay()
	fmt.Println("[struct] manager pay:", pay)
	employee.Display()

}

type PayRole interface {
	Display()
	CalPay() float64
}

type Employee struct {
	name   string
	salary float32
	bonus  float32
}

func (m *Employee) CalPay() float32 {
	return m.salary + m.bonus
}

func (m *Employee) Display() {
	fmt.Println("[struct] manager name:", m.name)
}

/*
  Go Uses I Table to implement interface

*/
