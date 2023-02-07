package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Unbuffered channel
	/*
		c := make(chan int)
		c <- 1
		fmt.Println(<-c)
	*/

	// Buffered channel
	/*
		c := make(chan int, 1)
		c <- 1
		fmt.Println(<-c)
	*/
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Ended %d\n", i)
}

/*

	Usando las Goroutines y el Waitgroup se tardo solo 2 segundos
	mientras que sin ellas se tardo 20 segundos xD

*/
