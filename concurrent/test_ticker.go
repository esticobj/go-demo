package main

import (
	"fmt"
	"time"
)

func test1() {
	t := time.NewTicker(time.Second * 2)
	i := 0
	for _ = range t.C {
		if i == 5 {
			t.Stop()
			break
		}
		fmt.Println("run...", i)
		i++
	}
}

func test2() {
	var ch = make(chan int)
	t := time.NewTicker(time.Second)
	go func() {
		for _ = range t.C {
			select {
			case ch <- 1:
			case ch <- 2:
			case ch <- 3:
			}
		}
	}()
	go func() {
		sum := 0
		for v := range ch {
			fmt.Printf("v: %v\n", v)
			sum += v
			if sum >= 10 {
				break
			}
		}
	}()
	time.Sleep(time.Second * 20)

}

func main() {
	// test1()
	test2()
}
