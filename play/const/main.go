package main

import "fmt"

const (
	// A is const
	A int = iota
	// B is const
	B
	// C is const
	C
)

func main() {
	fmt.Println("A", A)
	fmt.Println("B", B)
	fmt.Println("C", C)
}
