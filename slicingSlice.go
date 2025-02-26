package main

import "fmt"

func main() {
	var x = []int{23, 76, 45, 76, 78, 43, 34, 23, 12}
	fmt.Println("OG array: ", x)
	y := x[2:5]
	fmt.Println("Slicing array: ", y)
	x[2] = 32
	x[3] = 65
	x[4] = 99
	//x[5] = 11
	fmt.Println("New Array: ", x)
}
