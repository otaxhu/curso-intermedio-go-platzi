package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	id int
}

type fullTimeEmployee struct {
	person
	employee
}

func newFullTimeEmployee(name string, age int, id int) *fullTimeEmployee {
	return &fullTimeEmployee{
		person{
			name: name,
			age:  age,
		}, employee{
			id: id,
		},
	}
}

func main() {
	ftEmployee := fullTimeEmployee{}
	fmt.Println(ftEmployee)
	ftEmployee.name = "oscar"
	ftEmployee.age = 17
	ftEmployee.id = 2
	fmt.Println(ftEmployee)
	ftEmployee2 := newFullTimeEmployee("pepito", 14, 7)
	fmt.Println(*ftEmployee2)
}
