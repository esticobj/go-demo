package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("os.Environ(): %v\n", os.Environ())
	fmt.Printf("os.Getenv(\"GOPATH\"): %v\n", os.Getenv("GOPATH"))
	fmt.Printf("os.Getenv(\"javahome\"): %v\n", os.Getenv("javahome"))
	s, b := os.LookupEnv("javahome")
	if b {
		fmt.Printf("s: %v\n", s)
	} else {
		fmt.Println("不存在")
	}
	os.Setenv("javahome", "D:/")
	fmt.Printf("os.Getenv(\"javahome\"): %v\n", os.Getenv("javahome"))
}
