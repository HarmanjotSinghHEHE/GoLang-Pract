package main

import "fmt"

func main() {
	// Create a slice with 3 elements
	var s = []int{1, 2, 3}
	fmt.Println("Initial Slice:", s)
	fmt.Println("Length:", len(s), "Capacity:", cap(s))

	// Append elements to the slice
	s = append(s, 4)
	fmt.Println("\nAfter Appending 4:")
	fmt.Println("Slice:", s)
	fmt.Println("Length:", len(s), "Capacity:", cap(s))

	s = append(s, 5)
	fmt.Println("\nAfter Appending 5:")
	fmt.Println("Slice:", s)
	fmt.Println("Length:", len(s), "Capacity:", cap(s))

	s = append(s, 6)
	fmt.Println("\nAfter Appending 6:")
	fmt.Println("Slice:", s)
	fmt.Println("Length:", len(s), "Capacity:", cap(s))

	s = append(s, 7)
	fmt.Println("After apending 7:")
	fmt.Println("Slice: ", s)
	fmt.Println("Length:", len(s), "Capacity:", cap(s))
}
