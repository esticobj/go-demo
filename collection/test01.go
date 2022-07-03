package main

import "fmt"

func main() {
	var map1 map[string]int
	type s struct {
		name string
		age  int
	}
	fmt.Printf("map1: %v\n", map1 == nil)
	var x []int = make([]int, 3)
	x[0] = 1
	fmt.Printf("x: %v\n", x[1])
}
