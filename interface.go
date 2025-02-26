package main
import "fmt"

func main{
var i interface{} = "hello"

s := i.(string) // type assertion
fmt.Println(s)

s, ok := i.(string) // safe type assertion
if ok {
    fmt.Println(s)
} else {
    fmt.Println("not a string")
}
}