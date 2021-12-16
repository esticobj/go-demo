package main

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func cal(base int) (func(a int) int, func(a int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub
}

//闭包=函数+引用环境
func main() {
	f := add()
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(20): %v\n", f(20))
	g := add()
	fmt.Printf("g(10): %v\n", g(10))
	fmt.Printf("g(20): %v\n", g(20))

	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(20): %v\n", f(20))

	f1, f2 := cal(10)
	fmt.Printf("f1(1): %v\n", f1(1))
	fmt.Printf("f2(1): %v\n", f2(1))
}
