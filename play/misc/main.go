package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Printf("%v %#v\n", a, &a)
	recvarg(a)

	var b bool
	var c, d string = "c", "d"
	var e float32
	var f float64
	f = 1.11
	var g = f
	fmt.Println(a, b, c, d, e, f, g)
}

func recvarg(n int) {
	fmt.Printf("recvarg %v %#v\n", n, &n)
}
