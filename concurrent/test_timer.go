package main

import (
	"fmt"
	"time"
)

func test1() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())
	t1 := <-timer.C
	// <-timer.C //如果只是为了等待可以放弃返回值
	fmt.Printf("t1: %v\n", t1)
}

func test2() {
	<-time.After(time.Second * 3)
	fmt.Println("end...")
}

func test3() {
	timer := time.NewTimer(time.Second * 10)
	go func() {
		<-timer.C
		fmt.Println("end...") // 停止掉timer，该语句不会执行
	}()
	<-time.After(time.Second * 2)
	b := timer.Stop() // 停止timer
	if b {
		fmt.Println("stop success")
		<-time.After(time.Second * 5)
	}
}

func test4() {
	timer := time.NewTimer(time.Second * 10)
	go func() {
		<-timer.C
		fmt.Println("end...")
	}()
	<-time.After(time.Second * 2)
	b := timer.Reset(time.Second * 3) //重新设置timer的时间
	if b {
		fmt.Println("reset success")
		<-time.After(time.Second * 5)
	}
}
func main() {
	test1()
	test2()
	test3()
	test4()
}
