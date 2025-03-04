// reversing a string concurrntly
package main

import (
	"fmt"
	"sync"
)

func reverseStr(r rune, ch chan<- rune, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- r
}

func addingreverse(ch <-chan rune, length int) {
	runes := make([]rune, length)
	for i := length - 1; i >= 0; i-- {
		runes[i] = <-ch
	}
	reversed := string(runes) // Convert back to string
	fmt.Println("Reversed string:", reversed)
}

func main() {
	str := "Hello I am harman"
	ch := make(chan rune, len(str))
	var wg sync.WaitGroup

	for _, val := range str {
		wg.Add(1)
		reverseStr(val, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	addingreverse(ch, len(str))
}
