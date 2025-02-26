package main

import "fmt"

func main() {
	nslice := make([]int, 10)
	fmt.Println("	Capacity Initiak", cap(nslice))
	nslice = append(nslice, 1)
	fmt.Println("	Capacity 1 element", cap(nslice))
	nslice = append(nslice, 2)
	fmt.Println("	Capacity 2 element", cap(nslice))
	nslice = append(nslice, 3)
	fmt.Println("	Capacity 3 element", cap(nslice))
	nslice = append(nslice, 4)
	fmt.Println("	Capacity 4 element", cap(nslice))
	nslice = append(nslice, 5)
	fmt.Println("	Capacity 5 element", cap(nslice))
}
