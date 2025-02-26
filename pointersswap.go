package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	x, y := 5, 10
	fmt.Println("Before Swapping ", x, y)
	swap(&x, &y)
	fmt.Println("\nAfter Swapping ", x, y)

}
