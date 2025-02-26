package main

import "fmt"

func main() {
	nslice := make([]int, 10, 20)
	nslice = append(nslice, 1)
	nslice = append(nslice, 2)
	nslice = append(nslice, 3)
	nslice = append(nslice, 4)
	mslice := make([]int, 20, 20)

	xslice := copy(mslice, nslice)
	fmt.Println(xslice)
	fmt.Println(mslice)
}
