package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 2; i++ {
		go readChannel(i, ch)
	}
	time.Sleep(time.Second)
	ch <- "hello a"
	ch <- "hello b"
	ch <- "hello c"
	fmt.Println("all sent")
	for {
		time.Sleep(time.Second)
	}
}

func readChannel(i int, ch chan string) {
	fmt.Println(i, "reading")
	s := <-ch
	fmt.Println(i, "read s", s)
}
