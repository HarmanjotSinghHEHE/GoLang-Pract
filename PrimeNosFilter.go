// finding out a primenumber
package main

import (
	"fmt"
	"sync"
)

func findPrime(num int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num <= 1 {
		return
	}
	for i := 2; i*i < num; i++ {
		if num%i == 0 {
			return
		}
	}
	ch <- num
}

func addtolist(ch <-chan int) []int {
	var list []int

	for num := range ch {
		list = append(list, num)
	}
	return list
}

func main() {
	nums := []int{2, 4, 2, 32, 43, 3, 3, 54, 2, 32, 31, 21}
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go findPrime(num, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	//addtolist(ch)
	primes := addtolist(ch)
	fmt.Println("Prime numbers:", primes)
}
