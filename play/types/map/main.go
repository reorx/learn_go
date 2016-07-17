package main

import (
	"fmt"
	"strings"
)

type A struct {
	x int
	Y int
}

type B map[string]int

func main() {
	a := make(map[int]string)
	a[9] = "p"
	a[8] = "oo"
	fmt.Println(a)

	var b = map[string]A{
		"wtf": A{
			1, 2,
		},
		"foo": {2, 3},
	}
	fmt.Println(b)
	bv := b["wtf"]
	fmt.Println(bv)
	delete(b, "foo")
	bvx, bvy := b["foo"]
	fmt.Println("not in ", bvx, bvy)
	fmt.Println("not in ", b["qua"])

	c := make(map[string]int)
	fmt.Println(c)

	d := make(B)
	fmt.Println(d)

	// fmt.Println(A.Y)

	// panic: assignment to entry in nil map
	// var c map[string]int
	// c["a"] = 1
	// fmt.Println(c)
}

func WordCount1(s string) map[string]int {
	c := make(map[string]int)
	for _, word := range strings.Fields(s) {
		v, isin := c[word]
		if isin {
			c[word] = v + 1
		} else {
			c[word] = 1
		}
	}
	return c
}

func WordCount2(s string) map[string]int {
	c := make(map[string]int)
	for _, word := range strings.Fields(s) {
		c[word] += 1
	}
	return c
}
