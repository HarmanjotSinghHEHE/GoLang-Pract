package main

import (
	"fmt"
	"time"
)

func senddata(ch chan string) {
	ch <- "hello"
	time.sleep(1 * time.second)
	ch <- "world"
}

func main() {
	ch := make(chan string)
	go senddata(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Main function finished")
}
