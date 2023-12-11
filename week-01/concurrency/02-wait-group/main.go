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

	wg.Add(1)
	go func() {
		defer wg.Done()
		printLetters()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		printNumbers()
	}()

	wg.Wait()
}

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	var letters = "abcdefghij"
	for idx := range letters {
		fmt.Println(string(letters[idx]))
	}
}
