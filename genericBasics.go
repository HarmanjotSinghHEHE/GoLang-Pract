package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max(3, 5))
	fmt.Println(max(3.2342, 5.232))
	fmt.Println(max("apple", "banana"))
}
