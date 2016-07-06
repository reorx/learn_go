package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    fmt.Println("math main")
    fmt.Println("rand number", rand.Intn(100))
    fmt.Println("sqrt", math.Sqrt(9))
    fmt.Println("pie", math.Pi)

    g, h := 3, 4
    i := math.Sqrt(float64(g*g + h*h + 1))
    j := uint(i)
    fmt.Println(i, j)

    const (
        big = 1 << 100
        small = big >> 99
    )
    bigf := big * 0.1
    fmt.Println("bit shift", bigf, small)
}
