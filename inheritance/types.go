package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// Composicion sobre herencia o struct de structs
type FullTimeEmployee struct {
	Person
	Employee
	endDate int
}

func (ft FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (t TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}
