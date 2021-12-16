package main

import (
	"fmt"
	"os"

	"example.com/m/util"
	_ "example.com/m/util2" // 只调用包里的init()函数，不调用其他函数
)

func cal(a, b int, c string) (res int, str string) {
	return a + b, c[2:5]
	// 两种返回方式均可
	// res = a + b
	// str = c
	// return
}

func fun(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

// str会自动分配在堆上
func escape() *string {
	str := "Hello"
	return &str
}

func main() {
	res, str := cal(10, 20, "Hello World")
	fmt.Println("res:", res, "str:", str)
	fmt.Println("os.Args:", os.Args)
	fun(1, 2, 3)
	fun(4, 5, 6, 7, 8)
	// go不允许函数嵌套定义，但允许匿名函数
	//立即执行函数
	rs := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}(1, 2)
	fmt.Printf("rs: %v\n", rs)

	p := escape()
	fmt.Println("*p", *p)

	fmt.Println("res:", util.Add(1, 2))

	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "-c":
		fmt.Println("ccc")
		// 自动会添加break，如果不需要可加fallthrough
	case "-d":
		fmt.Println("ddd")
		fmt.Println("123")
		fallthrough
	default:
		fmt.Println("default")
	}

}
