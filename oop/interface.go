package main

import "fmt"

func main() {
	var i, j interface{} //定义接口类型
	i = 123
	j = "fgdh"
	fmt.Println("i:", i, "j:", j)
	iValue, ok := i.(bool) // .(int)是类型转换函数
	if ok {
		fmt.Println("i是bool:", iValue)
	} else {
		fmt.Println("i不是bool")
	}
	f(123)
	f(Person{"aaa", 123, "男"})
	f(true)
	f(int8(126))
}

type Person struct {
	name   string
	age    int
	gender string
}

func f(param interface{}) {
	switch param.(type) { //.(type)只能在switch中使用
	case int:
		fmt.Println("int", param.(int))
	case string:
		fmt.Println("int", param.(int))

	case Person:
		fmt.Println("Person,name:", param.(Person).name)
	default:
		fmt.Println("other type")
	}
}
