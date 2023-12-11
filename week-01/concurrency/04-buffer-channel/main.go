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
	var intch = make(chan int, 5) // using buffer

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

/*
	Note :
		- Penggunaan buffer pada producer membuat producer mengumpulkan 
		  nilainya di dalam buffer jika consumer belum siap menerima.

		  Hingga buffer penuh kemudian operasi akan ter block pada line 39
		  jika consumer masih belum siap.

		  penggunaan buffer pada producer ini berguna untuk  
		  meng-syncronisasikan fast producer dan slow consumer.
*/
