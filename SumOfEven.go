// printing sum of even number , use goRoutines to identify even numbers
package main

import (
	"fmt"
	"sync"
)

func checkEven(num int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	if num%2 == 0 {
		ch <- num
	}
}

func aggregateEven(ch <-chan int, len int) {
	sum := 0
	for i := 0; i < len; i++ {
		ev := <-ch
		sum = sum + ev
	}
	fmt.Println(sum)
}

func main() {
	nums := []int{12, 44, 46, 48, 90, 21, 23, 45, 465, 67, 45, 3, 4, 67, 78, 89, 89, 67, 88}
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go checkEven(num, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	go aggregateEven(ch, len(nums))
}
