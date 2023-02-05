package main

import "fmt"

type employee struct {
	id         int
	name       string
	isVacation bool
}

//		Constructor como funcion
func newEmployee(id int, name string, isVacation bool) *employee {
	return &employee{
		id:         id,
		name:       name,
		isVacation: isVacation,
	}
}

func main() {
	// Constructor sin inicializar (con zero values por defecto)
	e := employee{}
	fmt.Println(e)

	// Constructor inicializado asignando valores a las propiedades
	e2 := employee{
		// Values
		id:         2,
		name:       "constructor 2",
		isVacation: true,
	}
	fmt.Println(e2)

	// Constructor con la palabra reservada "new" (con zero values por defecto)
	e3 := new(employee)
	fmt.Println(*e3)

	// Instanciando funcion constructora
	e4 := newEmployee(9, "otro constructor", true)
	fmt.Println(*e4)
}
