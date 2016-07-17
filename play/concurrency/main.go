package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func mainSay() {
	fmt.Println("---------mainSay----------")
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	fmt.Println("call sum", s[0])
	sum := 0
	for _, v := range s {
		fmt.Println(" sum", s[0])
		sum += v
	}
	c <- sum // send sum to c
	fmt.Println("sent chan", s[0])
}

func mainSum() {
	fmt.Println("---------mainSun----------")
	s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int, 10)
	c := make(chan int)
	fmt.Println("before go go")
	go sum(s[:len(s)/2], c)
	fmt.Println("go one")
	go sum(s[len(s)/2:], c)
	fmt.Println("go two")
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func mainSend() {
	fmt.Println("---------mainSend----------")
	// ch := make(chan int)
	ch := make(chan int, 4)
	go send(ch, 1)
	go send(ch, 2)
	go send(ch, 3)
	// go send(ch, 4)
	time.Sleep(100 * time.Millisecond)
	recv(ch)
	recv(ch)
	recv(ch)
	// recv(ch)
}

func send(c chan int, n int) {
	fmt.Println("sent", n, "cap", cap(c), "len", len(c))
	c <- n
}

func recv(c chan int) {
	v, err := <-c
	if !err {
		fmt.Println(err)
	}
	fmt.Println("recvd", v, "cap", cap(c), "len", len(c))
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func mainSelect() {
	fmt.Println("---------mainSelect----------")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// // Walk walks the tree t sending all values
// // from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int) {
// 	walk(t, ch)
// 	close(ch)
// }
//
// func walk(t *tree.Tree, ch chan int) {
// 	if t == nil {
// 		//fmt.Println("no tree")
// 		return
// 	}
// 	ch <- t.Value
// 	walk(t.Left, ch)
// 	walk(t.Right, ch)
// }
//
// func mainWalk() {
// 	fmt.Println("---------mainWalk----------")
// 	ch := make(chan int, 10)
// 	go Walk(tree.New(1), ch)
// 	for v := range ch {
// 		fmt.Println("v", v)
// 	}
// }

func main() {
	mainSay()
	mainSum()
	time.Sleep(100 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)
	mainSend()
	time.Sleep(100 * time.Millisecond)
	mainSelect()
}
