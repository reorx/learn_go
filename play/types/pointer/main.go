package main

import (
	"fmt"
)

func main() {
	fmt.Println("play pointer")

	var a *int
	b := 1
	a = &b
	fmt.Println("a, &a, *a, b, &b:", a, &a, *a, b, &b)

	m := 2
	fmt.Println(*&m)

	// b -> c -> d -> e
	c := &b
	d := &c
	e := &d
	fmt.Println("b, c, d, e:", b, c, d, e)
	fmt.Println(*e, **e, ***e)
	fmt.Println("*e == d", *e == d)
	// Impossible, while the pointer goes deep, type would be **int or ***int depends on level
	// *d = e
	// fmt.Println("set *d", d, *d)

	// w -> x -> y -> z
	w := 2
	x := &w
	y := &x
	z := &y
	fmt.Println("z and it's real value", z, ***z)
	z = e
	fmt.Println("z after sign", z, ***z)
	***z = 3
	fmt.Println("b after sign", b, ***e, w)

	f := "a str"
	Recvpointer1(&f)
	Recvpointer2(&f)

	g := A{3}
	g.Point()
	(&g).Point()
	gp := &g
	gp.x = 4
	// &g.Point()
	gp.Point()

	var h interface{}
	h = &g
	fmt.Printf("h %v %#v %T\n", h, h, h)

	var i *interface{}
	// i = &g
	i = &h
	fmt.Printf("i %v %#v %T\n", i, i, i)

	var j int = 1
	fmt.Println("j", j)
	k := &j
	*k = 2
	fmt.Println("j", j)
}

func Recvpointer1(p *string) {
	fmt.Printf("p %v %T %#v\n", p, p, p)
}

type stringp *string

func Recvpointer2(p stringp) {
	fmt.Printf("p %v %T %#v\n", p, p, p)
}

func Recvpointer3(p string) {
	fmt.Printf("p %v %T %#v\n", p, p, p)
}

type A struct {
	x int
}

func (a *A) Point() {
	fmt.Println("A.x", a.x)
}
