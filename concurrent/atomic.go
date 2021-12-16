package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var a int32 = 100

func add() {
	atomic.AddInt32(&a, 1)
}
func sub() {
	atomic.AddInt32(&a, -1)
}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("atomic.LoadInt32(&a): %v\n", atomic.LoadInt32(&a))
	atomic.StoreInt32(&a, 200)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("atomic.LoadInt32(&a): %v\n", atomic.LoadInt32(&a))
	b := atomic.CompareAndSwapInt32(&a, 201, 205)
	if b {
		fmt.Println("修改成功, a = ", a)
	} else {
		fmt.Println("修改失败, a = ", a)
	}
}
