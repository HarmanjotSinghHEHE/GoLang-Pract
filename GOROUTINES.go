package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Start a goroutine
	go printNumbers() // Start another goroutine

	// Wait for goroutines to finish (in a real program, use sync.WaitGroup or channels)
	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
}
