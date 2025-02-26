package main

import "fmt"

func main() {
    fmt.Println("Enter the Number for which you want to Print table till 10\n")
    var num int64
    var i int64 // Declare i explicitly as int64
    fmt.Scanln(&num)
    var prod int64 = 0
    for i = 1; i <= 10; i++ { // i is of type int64 now
        prod = num * i
        fmt.Println(num, " * ", i, " = ", prod)
    }
}
