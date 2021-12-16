package main

import "fmt"

func main() {
	str := `
		Hello World
	`
	fmt.Print("str=", str, "length=", len(str))
	name := "gsfgfsg"
	for i := 1; i < len(name); i++ {
		fmt.Println("i:", i, ",char:", name[i])
	}
}
