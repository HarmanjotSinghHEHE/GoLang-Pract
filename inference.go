package main

import "fmt"

// now on this wer are declaring a T any generic , that have a slice S and this funcytions iterates it over printing the value as for key we have given __
func ps[t any](s []t) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	ps([]int{1, 2, 3, 4})
	ps([]string{"a", "a", "a", "a", "a"})
}
