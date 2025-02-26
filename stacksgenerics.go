package main

import (
    "fmt"
)

// Stack is a generic LIFO (Last In, First Out) data structure.
type Stack[T any] struct {
    items []T
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
func (s *Stack[T]) Pop() T {
    if len(s.items) == 0 {
        panic("stack is empty")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func main() {
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    fmt.Println(intStack.Pop()) // 2
    fmt.Println(intStack.Pop()) // 1

    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
    fmt.Println(stringStack.Pop()) // "world"
    fmt.Println(stringStack.Pop()) // "hello"
}
