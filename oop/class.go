package main

import "fmt"

func main() {
	person := Person{"fds", 12, "男"}
	person.eat()
	fmt.Println("person.age", person.age)
	person.sit()
	fmt.Println("person.age", person.age)
}

type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) eat() {
	p.age = 19 // 不会修改对象本身
	fmt.Println(p.name, "eat")
}
func (this *Person) sit() {
	this.age = 19 // 会修改对象本身
	fmt.Println(this.name, "sit")
}
