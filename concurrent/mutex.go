package main

import (
	"fmt"
	"sync"
	"time"
)

var a = 100
var mutex sync.Mutex

func add() {
	mutex.Lock()
	a++
	mutex.Unlock()
}
func sub() {
	mutex.Lock()
	a--
	mutex.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second)
	fmt.Printf("a: %v\n", a)
}
