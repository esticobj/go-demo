package main

import (
	"fmt"
	"os"
)

func main() {
	readFile("var.go")
	fmt.Println("main end")
}

func readFile(fileName string) {
	fpr, err := os.Open(fileName)
	// defer fpr.Close()
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	defer func() { //栈退出时执行defer，多个defer按照先入后出的原则执行
		//如果执行29行代码输出文件内容，打包成exe会执行defer，但使用go run没有执行，不知道什么原因
		fmt.Println("准备关闭文件")
		fpr.Close()
	}()
	defer fmt.Println("123")
	buf := make([]byte, 1024)
	n, _ := fpr.Read(buf)
	fmt.Println("读取文件的实际长度:", n)
	// fmt.Println("文件内容:", string(buf))
	fmt.Println("End")
}
