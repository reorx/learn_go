package math

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestPlus(t *testing.T) {
    fmt.Println("TestPlus")
    var x, y int
    x = 1
    y = 2
    z := Plus(x, y)
    assert.Equal(t, 3, z, "should be equal")
}
