package main

import "fmt"

func reverse(str string) {
	var bs []rune = []rune(str)
	//fmt.Println(len(bs))

	for i := len(bs) - 1; i >= 0; i-- {
		fmt.Print(string(bs[i]))
	}
}

func palin(str string) bool {
	var bs []rune = []rune(str)
	left := 0
	right := len(bs) - 1

	for left < right {
		if bs[left] != bs[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func max_no(arr []int) int {
	max_ele := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_ele {
			max_ele = arr[i]
		}
	}
	return max_ele
}

func main() {
	var x string = "NAMAN"
	xx := []int{1, 2, 32423, 213, 123, 35, 4356, 47, 678, 6}
	reverse(x)
	fmt.Println("\n", palin(x))
	fmt.Println(max_no(xx))
}
