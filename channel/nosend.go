package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	fmt.Println("start", ch)
	fmt.Println(<-ch)
	//ch <- "hello"
	for {
		time.Sleep(time.Second)
	}
}
