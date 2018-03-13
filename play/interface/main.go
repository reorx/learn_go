package main

import (
	"fmt"
	"reflect"
	"time"
)

// Axes interface type
type Axes interface {
	Linear() int
}

// Foo impl Axes
type Foo struct {
	x, y int
}

// Linear impl of Foo
func (f Foo) Linear() int {
	fmt.Println("got", f)
	return f.x * f.y
}

// String of Foo
// NOTE pointer is no right: http://stackoverflow.com/a/16978611/596206
// func (f *Foo) String() string {
func (f Foo) String() string {
	return fmt.Sprintf("<Foo x:%v y:%v>", f.x, f.y)
}

// Bar impl Axes
type Bar struct {
	x, y int
}

type Qua struct {
	x, y int
}

// Linear impl of Bar
func (b *Bar) Linear() int {
	if b == nil {
		fmt.Println("got nil")
		return 0
	}
	return b.x * b.y * 2
}

func main() {
	var _a Axes
	fmt.Printf("_a %#v %T %#v \n", _a, _a, &_a)

	a := Foo{2, 3}
	fmt.Println(a, a.Linear())

	var b Axes
	// b = a
	b = &Foo{3, 4}
	fmt.Println(b, b.Linear())
	fmt.Printf("b %#v %T %#v \n", b, b, &b)

	var c Axes = Foo{4, 5}
	fmt.Println(c.Linear())

	ca := c.(Foo)
	fmt.Printf("c %#v %T %#v \n", c, c, &c)
	fmt.Printf("ca %#v %T %#v \n", ca, ca, &ca)

	var d Foo
	var e Axes = d
	fmt.Println(e.Linear())

	var f *Bar
	fmt.Printf("f: %v %T\n", f, f)
	if f == nil {
		fmt.Println("f is nil")
	} else {
		fmt.Println("f is not nil")
	}
	// this will fail
	// var f Bar
	var g Axes = f
	fmt.Println("g.Linear", g.Linear())

	switchtype(1)
	switchtype("string")

	h := getaxes(9, 8)
	fmt.Printf("h %#v %T %#v \n", h, h, &h)

	j := geterror()
	fmt.Println("got this", j)

	fmt.Println(PositiveOrDie(11))
	fmt.Println(PositiveOrDie(-2))

	k := Foo{10, 20}
	recvAxes(k)
	// l := Qua{1, 2}
	// recvAxes(l)

	vb := reflect.ValueOf(b)
	fmt.Println("reflect", vb, vb.Type())
	fmt.Println("describe vb.Type()")
	describe(vb.Type())
}

func recvAxes(a Axes) {
	fmt.Printf("recv axes, %v %#v\n", a, &a)
	fmt.Println(a.Linear())
}

type Any interface{}

// func describe(i interface{}) {
func describe(i Any) {
	fmt.Printf("%#v %T %#v \n", i, i, &i)
}

func switchtype(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v, "type is int", i)
	default:
		fmt.Println(v, "type is wat", i)
	}
}

func getaxes(x, y int) Axes {
	f := Foo{x, y}
	return f
}

// Error for error interface
type Error struct {
	What string
	When time.Time
}

func (e *Error) Error() string {
	return fmt.Sprintf("<Error: %v at %v>", e.What, e.When)
}

func (e Error) String() string {
	return fmt.Sprintf("'Error: %v at %v''", e.What, e.When)
}

func geterror() error {
	_e := &Error{
		"something wrong",
		time.Now(),
	}
	fmt.Println("constructing error", *_e)
	return _e
}

// PositiveOrDie return positive number
func PositiveOrDie(n float64) (float64, error) {
	if n < 0 {
		return 0, NegativeFloatError(n)
	}
	return n, nil
}

// NegativeFloatError error
type NegativeFloatError float64

func (e NegativeFloatError) Error() string {
	// NOTE should not use %v, it will cause recursive calling
	// return fmt.Sprintf("negative float error: %v", e)
	// %f is OK
	// return fmt.Sprintf("negative float error: %f", e)
	return fmt.Sprintf("negative float error: %g", e)
}
