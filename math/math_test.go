package math

import (
    "testing"
    "fmt"
)


func TestPlus(t *testing.T) {
    fmt.Println("TestPlus")
    var x, y int
    x = 1
    y = 2
    z := Plus(x, y)
    if z != 3 {
        t.Error("Expect 3, got ", z)
    }
}
