package main

import (
	"fmt"
)

func main() {
	runAsync()
}

func runAsync() {
	var oddch, evench = make(chan int), make(chan int)
	var errch = make(chan error)
	var donech = make(chan struct{})

	go func() {
		for i := 1; i <= 25; i++ {
			if i <= 20 && i%2 != 0 {
				oddch <- i
				continue
			}

			if i <= 20 && i%2 == 0 {
				evench <- i
				continue
			}

			errch <- fmt.Errorf("error: number %d is greater than 20", i)

		}
		close(oddch)
		close(evench)
		close(errch)
		close(donech)
	}()

	for {
		select {

		case num, ok := <-oddch:
			if !ok {
				continue
			}
			fmt.Println("RECEIVED ODD  :", num)

		case num, ok := <-evench:
			if !ok {
				continue
			}
			fmt.Println("RECEIVED EVEN :", num)

		case err, ok := <-errch:
			if !ok {
				continue
			}
			fmt.Println(err)

		case <-donech:
			return
		}
	}
}
