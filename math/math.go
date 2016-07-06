package math

import (
    "fmt"
)

// Plus x + y
func Plus(x, y int) int {
    fmt.Println("Plus")
    return x + y
}

// Minus x - y
func Minus(x, y int) (z int) {
    z = x - y
    return
}


const delta = 1e-8

// Sqrt implements Newton's algorithm
func Sqrt(x float64) float64 {
	z := float64(1)
    var zn float64
	for {
		zn = z - (z*z - x) / (2*z)
		if abs(zn - z) < delta {
			return zn
		}
		z = zn
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
