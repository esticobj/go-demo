package main

import (
	"fmt"
	"log"
	"os"
)

var logger log.Logger

func init() {
	f, err := os.OpenFile("out.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	logger = *log.New(f, "MyLog ", log.Ldate|log.Ltime|log.Lshortfile)
}

func test1() {
	defer log.Print("close resource")
	log.Print("Hello", " ", "World")
	log.Println("Hello", "World") // 不需要空格分隔
	log.Printf("Hello, %s, %d", "World", 100)
	// log.Panic("panic") //执行defer
	log.Fatal("fatal") // 不执行defer
}

func test2() {
	defer logger.Print("close resource")
	logger.Print("Hello", " ", "World")
	logger.Println("Hello", "World") // 不需要空格分隔
	logger.Printf("Hello, %s, %d", "World", 100)
	// logger.Panic("panic") //执行defer
	logger.Fatal("fatal") // 不执行defer
}
func main() {
	// test1()
	test2()
}
