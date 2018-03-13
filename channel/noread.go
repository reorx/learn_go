package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go readChannel(ch)

	ch <- "hello a"
	ch <- "hello b"
	ch <- "hello c"
	fmt.Println("all sent")
}

func readChannel(ch chan string) {
	for {
		fmt.Println("waiting")
		s := <-ch
		fmt.Println("read s", s)
		time.Sleep(time.Duration(3 * time.Second))
		fmt.Println("next loop")
	}
}
