package main

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
}
