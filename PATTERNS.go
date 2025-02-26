package main

import "fmt"

func rhs() {
	for i := 1; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("8")
		}
		fmt.Println()
	}
}

func square() {
	for i := 1; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Print("8")
		}
		fmt.Println()
	}
}

func lhs() {
	for i := 5; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("8")
		}
		fmt.Println()
	}
}

func main() {
	rhs()
	square()
	lhs()
}
