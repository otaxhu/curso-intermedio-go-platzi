package main

import "fmt"

type employee struct {
	id   int
	name string
}

func (e *employee) setId(id int) {
	e.id = id
}

func (e *employee) setName(name string) {
	e.name = name
}

func (e *employee) getId() int {
	//fmt.Println(e.id)
	return e.id
}

func (e *employee) getName() string {
	//fmt.Println(e.name)
	return e.name
}

func main() {
	e := employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "name"
	fmt.Println(e)
	e.setId(2)
	e.setName("other name")
	fmt.Println(e)
	fmt.Println(e.getId())
	fmt.Println(e.getName())
}
