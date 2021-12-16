package main

import "fmt"

func main() {
	var s1, s2 ISay
	s1 = &PersonA{"11"} //接口只能赋值地址
	s2 = &PersonB{"22"}
	s1.say()
	s2.say()
}

type ISay interface {
	say()
}

type PersonA struct {
	name string
}

func (p *PersonA) say() { //方法签名一样，就自动实现了接口ISay
	fmt.Println("haha,I'm", p.name)
}

type PersonB struct {
	name string
}

func (p *PersonB) say() {
	fmt.Println("hehe,I'm", p.name)
}

//不支持重载
// func test(msg string) {
// 	return "string"
// }
// func test(msg int) {
// 	return "int"
// }
