package main

import (
	"fmt"
	"time"
)

//type MapMutex {
//}

func main() {
	amap := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}

	go func() {
		for _, v := range amap {
			fmt.Println("go1 range", v)
			//time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for _, v := range amap {
			fmt.Println("go2 range", v)
			//time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		delete(amap, 3)
	}()

	time.Sleep(10 * time.Second)
}
