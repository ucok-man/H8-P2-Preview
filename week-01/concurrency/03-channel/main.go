package main

import (
	"fmt"
	"sync"
)

func main() {
	runAsync()
}

func runAsync() {
	var wg sync.WaitGroup
	intch := produce()

	wg.Add(1)
	go func() {
		defer wg.Done()
		consume(intch)
	}()

	wg.Wait()
}

func consume(intch <-chan int) {
	for num := range intch {
		fmt.Println(num)
	}
}

func produce() <-chan int {
	var wg sync.WaitGroup
	var intch = make(chan int)

	wg.Add(1)
	go func(inch chan<- int) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			intch <- i
		}
	}(intch)

	go func() {
		wg.Wait()
		close(intch)
	}()

	return intch
}
