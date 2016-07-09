package main

import (
    "fmt"
)

func main() {
    fmt.Println("play pointer")

    var a *int
    b := 1
    a = &b
    fmt.Println("a, &a, *a, b, &b:", a, &a, *a, b, &b)

    // b -> c -> d -> e
    c := &b
    d := &c
    e := &d
    fmt.Println("b, c, d, e:", b, c, d, e)
    fmt.Println(*e, **e, ***e)
    fmt.Println("*e == d", *e == d)
    // Impossible, while the pointer goes deep, type would be **int or ***int depends on level
    // *d = e
    // fmt.Println("set *d", d, *d)

    // w -> x -> y -> z
    w := 2
    x := &w
    y := &x
    z := &y
    fmt.Println("z and it's real value", z, ***z)
    z = e
    fmt.Println("z after sign", z, ***z)
    ***z = 3
    fmt.Println("b after sign", b, ***e, w)
}
