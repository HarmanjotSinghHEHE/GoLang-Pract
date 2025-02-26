package main

import "fmt"

func main() {
	x := 10
	y := 10
	ptrx := &x
	ptry := &y
	fmt.Println("Are pointers the same?", ptrx == ptry)

	ptry = &x
	fmt.Println("Are pointers the same now?", ptrx == ptry) // true
}
