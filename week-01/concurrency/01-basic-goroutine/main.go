package main

import (
	"fmt"
	"time"
)

func main() {
	// runSync()
	runAsync()
}

func runSync() {
	printLetters()
	printNumbers()
}

func runAsync() {
	go printLetters()
	go printNumbers()
	time.Sleep(1 * time.Second)
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
