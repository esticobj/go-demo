package main

import "fmt"

func main() {
	fmt.Println("test")
	x, y := 213, 456
	fmt.Println("x:", x, "y:", y)
	x, y = y, x
	fmt.Println("x:", x, "y:", y)
	var name string = "tom"
	fmt.Println("name:", name)
}
