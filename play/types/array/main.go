package main

import "fmt"

func main() {
	var a [3]int
	// a[0] = 1
	// a[1] = 9
	fmt.Println(a, len(a), cap(a))
	a1 := a[:2]
	fmt.Println(a1, len(a1), cap(a1))

	b := [3]string{"a", "b"}
	fmt.Println(b)

	c := make([]int, 2)
	fmt.Println(c)

	var d []int
	fmt.Println(d, len(d), cap(d))
	for i := 0; i < 10; i++ {
		d = append(d, i)
		fmt.Println(d, len(d), cap(d))
	}

	var e []int
	fmt.Println(e == nil)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	for j, k := range board {
		fmt.Println(j, k)
	}

	l := []byte("abc")
	fmt.Println(l)

	n := string(97)
	fmt.Println(n)

	list := []string{"aaa", "bbb", "ccc"}
	for _, el := range list {
		fmt.Println("element", el)
	}
}

func Pic1(dx, dy int) [][]uint8 {
	var l [][]uint8
	for i := 0; i < dy; i++ {
		var ll []uint8
		for j := 0; j < dx; j++ {
			ll = append(ll, uint8(200+(i*j)))
		}
		l = append(l, ll)
	}

	return l
}

func Pic2(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			a[i][j] = uint8(i + j)
		}
	}
	return a
}
