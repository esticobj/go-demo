package main

import "fmt"

func main() {
	//定义变量组
	var (
		x = 1
		y int
	)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	//定义常量组模拟枚举
	const (
		MONDEY    = iota // 0
		TUESDAY   = iota // 1, iota会递增
		WENDESDAY        // 常量组没有赋值默认与上一行相同
		THIRSDAY
		M, N = iota, iota //4,4 同一行的iota不会递增
	)
	const (
		X = iota + 1 // 1 在一个新的常量组中iota会清0
		Y            //2
	)
	fmt.Printf("WENDESDAY: %v\n", WENDESDAY)
	fmt.Printf("THIRSDAY: %v\n", THIRSDAY)
	fmt.Println("M:", M)
	fmt.Println("N:", N)
	fmt.Println("X:", X)
	fmt.Println("Y:", Y)
}
