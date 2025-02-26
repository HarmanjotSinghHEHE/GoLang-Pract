package main

import (
    "fmt"
)

//declaring an interface   that actys as a constraint

type number interface {
    int | float64 // | string
}

func sum[T number](a, b T) T {
    return a / b
}

func main() {
    fmt.Println(sum(3, 5))       // Works with int
    fmt.Println(sum(3.14, 2.71)) // Works with float64
    //fmt.Println(sum("world", "hello"))
}
