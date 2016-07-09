package main

import (
    "fmt"
)

// A struct
type A struct {
    X int
    x, y int
}

func main() {
    fmt.Println("types: struct")
    a := A{1, 2, 3}
    a.X = 10
    fmt.Println("a", a)

    pa := &a
    pa.x = 20
    fmt.Println("a, pa", a, pa, *pa)

    a1 := A{x: 1, y: 2}
    fmt.Println("a1", a1)

    a2 := A{}
    fmt.Println("a2", a2)

    pa1 := &A{}
    pa1.X = 1
    fmt.Println("pa1", pa1)
}
