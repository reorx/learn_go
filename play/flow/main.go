package main

import (
    "fmt"
)

func main() {
    fmt.Println("flow")
    c := 0
    for i := 0; i < 5; i ++ {
        c++
    }
    fmt.Println("c += ->", c)

    for c > 1 {
        c--
        d := c
        fmt.Println("c", c, "d", d)
    }

    if c = c + 1; c == 2 {
        fmt.Println("c == 2")
    }

}
