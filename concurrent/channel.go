package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	var ch chan int
	ch = make(chan int, 3)
	go produce(ch) // 如果生产个数与消费个数不对等，并且其中一个在主线程运行，将会发生死锁
	go consume(ch)
	time.Sleep(5 * time.Second)

}

func produce(ch chan<- int) { // 只能写入的通道
	defer func() {
		fmt.Println("produce end")
		close(ch)
	}()
	for i := 1; i <= 5; i++ {
		fmt.Println("生产者生产了消息，id:", i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
}

func consume(ch <-chan int) { //只能读取的通道
	// for i := 0; i < 3; i++ {
	// 	time.Sleep(20 * time.Millisecond)
	// 	x := <-ch
	// 	fmt.Println("消费者消费了消息，id:", x)
	// }

	for x := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("消费者消费了消息，id:", x)
	}
	// for {
	// 	time.Sleep(500 * time.Millisecond)
	// 	x, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("消费者消费了消息，id:", x)
	// }
	fmt.Println("consume end")
}
