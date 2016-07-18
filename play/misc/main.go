package main

import (
	"fmt"
)

func main() {
	// var <name> <type>
	// <name> = <value>
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

	ln()

	// var <name> = <value>
	var h = 1.1
	describe(h)
	var i float32 = 1.1
	describe(i)
	var j = "str"
	describe(j)
}

func recvarg(n int) {
	fmt.Printf("recvarg %v %#v\n", n, &n)
}

func ln() {
	fmt.Println("-----------")
}

func describe(i interface{}) {
	fmt.Printf("%v %T\n", i, i)
}
