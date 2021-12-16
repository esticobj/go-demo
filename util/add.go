package util

import "fmt"

//方法名首字母大写可以被外部使用，小写不能
func Add(a, b int) int {
	//init() // 不允许手动调用init()
	return a + b
}

//init函数可实现包级别的初始化操作
//init函数不能手动调用
//导入包的时候会自动调用init函数，可以有多个init函数，调用顺序未定义
func init() {
	fmt.Println("add init1")
}

func init() {
	fmt.Println("add init2")
}
