package main

import "fmt"

// Funcion variadica que retorna la suma de los valores de un slice
func sum(values ...int) int {
	var total int
	for _, num := range values {
		total += num
	}
	return total
}

// Funcion variadica que imprime un slices de strings
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// Funcion con retorno de variables
func getValues(x int) (double, triple, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	numero := sum(1, 2, 3)
	fmt.Println(numero)
	printNames("sofia", "oscar", "luis")
	fmt.Println(getValues(3))
}
