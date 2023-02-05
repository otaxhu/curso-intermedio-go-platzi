package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	c := make(chan int)

	// La mejor forma de hacer una Goroutine, con una funcion anonima
	go func() {
		fmt.Println("start")
		time.Sleep(5 * time.Second)
		fmt.Println("end")
		c <- 1
	}()
	<-c
}
