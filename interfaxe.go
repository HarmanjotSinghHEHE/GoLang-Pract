package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type rectange struct {
	Width  float64
	Height float64
}

func (r rectange) Area() float64 {
	return r.Width * r.Height
}

func (r rectange) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type circle struct {
	Radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func PrintShapeDetails(s shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := rectange{Width: 10, Height: 5}
	circle := circle{Radius: 7}

	PrintShapeDetails(rect)
	PrintShapeDetails(circle)
}
