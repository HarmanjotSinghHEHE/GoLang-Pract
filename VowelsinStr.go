// find vowels in a string
package main

import (
	"fmt"
	"strings"
	"sync"
)

func findVowels(alp rune, ch chan<- rune, wg *sync.WaitGroup) {
	defer wg.Done()
	if strings.ContainsRune("aeiouAEIOU", alp) {
		ch <- alp
	}
	//fmt.Println("")
}

func findCount(ch <-chan rune, len int) {
	c := 0
	for range ch { // Iterate over the channel to count vowels
		c++
	}
	fmt.Println("Total vowels:", c)
}

func main() {
	str := []rune{'f', 'D', 'R', 'F', 'W', 'Q', 'A', 'D', 'a', 'c', 'g', 'h', 'i', 'k', 'n', 'v', 'f', 'r', 'w', 's', 'z', 'c', 'h', 'j', 'o', 'u'}
	ch := make(chan rune)
	var wg sync.WaitGroup

	for _, val := range str {
		wg.Add(1)
		go findVowels(val, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	findCount(ch, len(str))

}
