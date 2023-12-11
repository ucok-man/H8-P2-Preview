package main

import "fmt"

func main() {
	runAsync()
}

func runAsync() {
	var oddch, evench = make(chan int), make(chan int)
	var donech = make(chan struct{})

	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 != 0 {
				oddch <- i
				continue
			}

			if i%2 == 0 {
				evench <- i
				continue
			}
		}
		close(oddch)
		close(evench)
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

		case <-donech:
			return
		}
	}
}
