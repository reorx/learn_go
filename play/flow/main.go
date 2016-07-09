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

        switch c {
        case 3:
            fmt.Println("c is in case 3")
        case 1:
            fmt.Println("c is in case 1")
        default:
            fmt.Println("c is in default case")
        }
    }

    if c = c + 1; c == 2 {
        fmt.Println("c == 2")
    }

    switch {
        case c < 0:
            fmt.Println("c < 0")
        case c == 0:
            fmt.Println("c == 0")
        // case c > 0:
        //     fmt.Println("c > 0")
    }

    trydefer()
}

func trydefer() {
    x := 1
    defer fmt.Println("deferred x value (1st)", x)
    x++
    fmt.Println("x", x)
    defer fmt.Println("deferred x value (2nd)", x)
    x++
    fmt.Println("x", x)
}
