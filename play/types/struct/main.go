package main

import (
	"fmt"
	"time"
)

// A struct
type A struct {
	X    int
	x, y int
}

type B struct {
	A
}

func main() {
	fmt.Println("types: struct")
	a := A{1, 2, 3}
	a.X = 10
	fmt.Println("a", a)

	pa := &a
	pa.x = 20
	fmt.Println("a, pa", a, pa, *pa)

	a1 := A{x: 1, y: 2}
	fmt.Println("a1", a1)

	a2 := A{}
	fmt.Println("a2", a2)

	pa1 := &A{}
	pa1.X = 1
	fmt.Println("pa1", pa1)

	{
		fmt.Println("* copy")
		a := &A{1, 2, 3}
		fmt.Printf("a %v %p \n", a, a)
		b := *a
		// b.X = 10
		fmt.Println("b", b)
		c := &b
		c.X = 20
		fmt.Println("c", c)
		a = &b
		fmt.Printf("a %v %p \n", a, a)
	}
	{
		type C struct {
			C1 string
			B  bool
			I  int
			A  time.Time
		}

		c := C{}
		fmt.Println("c", c)
		var c1 C
		fmt.Printf("c1 %v %#v %T", c1, c1, c1)
	}
	{
		type T struct {
			A string
			B int
			C int
		}
		l := []*T{
			{"1", 2, 3},
		}
		fmt.Println("last")
		fmt.Println("l", l[0])
	}

	// b := B{1, 2, 3}
	// fmt.Println("b", b)
}
