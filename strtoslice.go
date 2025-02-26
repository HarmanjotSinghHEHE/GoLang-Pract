package main

import "fmt"

func main() {
	var s string = "Hello 😃"
	var bs []byte = []byte(s)
	var cs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(cs)

}
