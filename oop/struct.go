package main

import "fmt"

func main() {
	//类型定义（定义了新的类型,相当于C语言的typedef）
	type MyInt int
	var x MyInt
	x = 3
	fmt.Printf("x: %v,type: %T\n", x, x)
	// 类型别名(没有定义新的类型)
	type Integer = int
	var y Integer = 2
	fmt.Printf("y: %v,type: %T\n", y, y)
	//定义结构体，成员不需要逗号
	type Student struct {
		name  string
		age   int
		score float64
	}
	//结构体初始化，可以只初始化部分成员
	person := Student{
		name:  "jerry",
		age:   18,
		score: 98.5, //最后一个元素必须加逗号，否则必须与}同行
	}
	fmt.Println("person:", person)
	px := &person
	fmt.Println("name", person.name, px.name, (*px).name)
	//结构体初始化，必须按顺序初始化全部成员
	person = Student{"aaa", 19, 99}
	fmt.Println("person:", person)
	//使用new创建结构体变量
	person2 := new(Student)
	person2.name = "bbb"
	fmt.Println("person2:", *person2)
	//匿名结构体
	var tom struct {
		name string
		age  int
	}
	tom.name = "Tom"
	tom.age = 20
	fmt.Printf("tom: %v\n", tom)
}
