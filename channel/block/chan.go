package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Printf("channel %v \n", ch)

	go func() {
		for {
			v := <-ch
			fmt.Println("received", v)
			time.Sleep(3 * time.Second)
		}
	}()
	time.Sleep(1 * time.Second)

	select {
	case ch <- 1:
	default:
		fmt.Println("skip 1")
	}
	time.Sleep(1 * time.Second)

	close(ch)

	//ch <- 2
	//fmt.Println("2 sent")

	select {
	case ch <- 2:
		fmt.Println("2 sent")
	default:
		fmt.Println("skip 2")
	}

	a := 0
	select {
	case ch <- a:
		fmt.Println("sent message", a)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout ch <- a")
		//default:
		//	fmt.Println("no message sent")
	}

	//ch <- 1
	//fmt.Println("1 sent")
	//ch <- 2
	//fmt.Println("2 sent")
	time.Sleep(10)
}
