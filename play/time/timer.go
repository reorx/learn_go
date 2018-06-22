package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	main1()
}

func main1() {
	n := time.Now()
	fmt.Println("testing timer")
	var t *time.Timer
	t = time.AfterFunc(2*time.Second, func() {
		fmt.Println("callback triggered ", time.Now().Sub(n))
		//fmt.Println("  reset ", t.Reset(time.Second))
	})
	fmt.Sprint(t)

	time.Sleep(time.Second * 3)

	fmt.Println("reset timer ", time.Now().Sub(n))
	t.Reset(2 * time.Second)
	fmt.Println("reset timer end")

	time.Sleep(time.Second * 1)

	fmt.Println("reset timer x2")
	t.Reset(2 * time.Second)
	fmt.Println("reset timer end")

	time.Sleep(3 * time.Second)

	fmt.Printf("exit, duration %s\n", time.Now().Sub(n))
}

func main2() {
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func waitTimer(t *time.Timer, f func()) {
	for {
		select {
		case <-t.C:
			fmt.Println("timer stopped ", t)
			if f != nil {
				f()
			}
		default:
			time.Sleep(time.Millisecond * 200)
			fmt.Println("waiting")
		}
	}
	/*
		<-t.C
		fmt.Println("timer stopped ", t)
	*/
}

func main3() {
	t := time.NewTimer(time.Second)
	go waitTimer(t, func() {
		t.Reset(time.Second)
	})

	t1 := time.AfterFunc(2*time.Second, func() {
		fmt.Println("callback triggered")
		//fmt.Println("  reset ", t.Reset(time.Second))
	})
	//go waitTimer(t1, nil)
	time.Sleep(3 * time.Second)
	t1.Reset(time.Second)

	time.Sleep(3 * time.Second)
}
