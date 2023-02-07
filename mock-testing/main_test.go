package main

import "testing"

func TestGetFullTimeEmployee(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "31655193",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{id: 1, position: "CEO"}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{DNI: "31655193", name: "oscar", age: 17}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					age:  17,
					DNI:  "31655193",
					name: "oscar",
				},
				Employee: Employee{
					id:       1,
					position: "CEO",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI
	for _, item := range table {
		item.mockFunc()
		ftEmployee, err := GetFulltimeEmployeeById(item.id, item.dni)
		if err != nil {
			t.Errorf("Error when getting full time employee")
		}
		if ftEmployee.age != item.expectedEmployee.age {
			t.Errorf("Error with the age, got %d expected %d", ftEmployee.age, item.expectedEmployee.age)
		}
		if ftEmployee.id != item.expectedEmployee.id {
			t.Errorf("Error with the id, got %d expected %d", ftEmployee.id, item.expectedEmployee.id)
		}
		if ftEmployee.name != item.expectedEmployee.name {
			t.Errorf("Error with the name, got %s expected %s", ftEmployee.name, item.expectedEmployee.name)
		}
		if ftEmployee.position != item.expectedEmployee.position {
			t.Errorf("Error with the position, got %s expected %s", ftEmployee.position, item.expectedEmployee.DNI)
		}
		if ftEmployee.DNI != item.expectedEmployee.DNI {
			t.Errorf("Error with the DNI, got %s expected %s", ftEmployee.DNI, item.expectedEmployee.DNI)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
