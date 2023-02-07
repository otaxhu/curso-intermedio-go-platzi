package main

type Person struct {
	DNI  string
	name string
	age  int
}

type Employee struct {
	id       int
	position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	return Employee{}, nil
}

func GetFulltimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	e, err1 := GetEmployeeById(id)
	if err1 != nil {
		return ftEmployee, err1
	}
	ftEmployee.Employee = e
	p, err2 := GetPersonByDNI(dni)
	if err2 != nil {
		return ftEmployee, err2
	}
	ftEmployee.Person = p
	return ftEmployee, nil

}
