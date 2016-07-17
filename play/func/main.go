package main

import (
	"fmt"
)

func callfunc(fn func(int, int) int) int {
	return fn(1, 2)
}

func adder(base int) func(int) int {
	return func(x int) int {
		base += x
		return base
	}
}

func main() {
	funca := func(x, y int) int {
		return x + y
	}

	// No `func` definition inside function
	// func localcall(fn func(int, int) int) int {
	localcall := func(fn func(int, int) int) int {
		return fn(3, 2)
	}

	fmt.Println(callfunc(funca))
	fmt.Println(localcall(funca))

	adda := adder(10)
	addb := adder(100)

	fmt.Println(adda(1))
	fmt.Println(addb(1))
	fmt.Println(adda(2))
	fmt.Println(addb(2))

	funcb := func() {
		fmt.Println("run funcb")
	}
	funcb()

	var funcd func()
	funcd = func() {
		fmt.Println("run funcd")
	}
	funcd()

	var funcc func(int) int
	funcc = func(a int) int {
		return a + 1
	}
	fmt.Println(funcc(1))

	defer func() {
		fmt.Println("declare to run later")
	}()

	func() {
		fmt.Println("declare and run")
	}()
}

func fibonacci() func() int {
	i := 0
	n1 := 0
	n2 := 1
	n3 := n1 + n2
	return func() int {
		var rv int
		switch i {
		case 0:
			rv = 0
		case 1:
			rv = 1
		default:
			rv = n3
			n1, n2 = n2, n3
			n3 = n1 + n2
		}
		i++
		return rv
	}
}
