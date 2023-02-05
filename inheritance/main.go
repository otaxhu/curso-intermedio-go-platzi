package main

import "fmt"

func main() {
	ftEmployee := FullTimeEmployee{}

	fmt.Println(ftEmployee)

	ftEmployee.name = "oscar"
	ftEmployee.age = 17
	ftEmployee.id = 2

	fmt.Println(ftEmployee)

	//ftEmployee2 := NewFullTimeEmployee("pepito", 14, 7)

	//fmt.Println(*ftEmployee2)

	//fmt.Println(ftEmployee2.age)

	tEmployee := TemporaryEmployee{}

	getMessage(tEmployee)
	getMessage(ftEmployee)
}
