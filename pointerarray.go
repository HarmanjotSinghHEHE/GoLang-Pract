package main

import "fmt"

//pointer to an array

func updatearray(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 5
	}
}

func updateslice(slice *[]int) {

}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before Updating ", arr, cap(arr))

	updatearray(&arr)
	arr[4] = 10
	fmt.Println("\nAfter Updating ", arr, cap(arr))

}
