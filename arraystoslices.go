package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 10)
	num := copy(y, x)
	fmt.Println(y, num)
}
