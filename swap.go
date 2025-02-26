package main

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x
}
func main() {
	var x int = 4
	var y int = 5

	a, b := swap(x, y)
	fmt.Println(a, b)
}
