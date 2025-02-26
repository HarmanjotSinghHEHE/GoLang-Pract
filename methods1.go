package main

import "fmt"

type si struct {
	P int
	R int
	T int
}

func (x si) calc() int {
	return (x.P * x.R * x.T) / 100
}

func main() {
	y := si{1200, 12, 5}
	fmt.Println(y.calc())
}
