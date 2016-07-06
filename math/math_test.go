package math

import (
    "fmt"
    "testing"
    "math"
    "github.com/stretchr/testify/assert"
)


func TestPlus(t *testing.T) {
    var x, y int
    x = 1
    y = 2
    z := Plus(x, y)
    assert.Equal(t, 3, z)
}


func TestMinus(t *testing.T) {
    x, y := 1, 2
    z := Minus(x, y)
    assert.Equal(t, -1, z)
}


func TestSqrt(t *testing.T) {
    x := 2.
    s := Sqrt(x)
    expected := math.Sqrt(x)
    fmt.Println("result", s, "expected", expected)
    // OMG it's equal!
    assert.Equal(t, expected, s)
}
