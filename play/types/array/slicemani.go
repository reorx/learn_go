package main

import "fmt"

func describe(b []int) {
	fmt.Println(b, "len", len(b), "cap", cap(b))
}

func getlen(b []int) int {
	fmt.Println("call getlen")
	return len(b)
}

func manipulate(b []int) {
	bl := getlen(b)
	var dft int
	// NOTE getlen will be called in every loop
	// for i := 0; i < getlen(b); i++ {
	for i := 0; i < bl; i++ {
		b[i] = dft
	}
}

func main() {
	b := make([]int, 10)
	for i := 0; i < 5; i++ {
		b[i] = i + 1
		describe(b)
	}

	manipulate(b)
	describe(b)

	var c byte
	c = 'A'
	fmt.Printf("c %v %T \n", c, c)
	c += 13
	fmt.Printf("c %v %T \n", c, c)

	d := 'B'
	fmt.Printf("d %v %T \n", d, d)
}
