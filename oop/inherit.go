package main

import "fmt"

func main() {
	s1 := Student{78, Human{"eras", 18}}
	fmt.Println("s1:", s1)
	fmt.Println("s1.name", s1.name)
	s1.eat()
}

type Human struct {
	name string
	age  int
}

func (p Human) eat() {
	fmt.Println(p.name, "eat")
}

type Student struct {
	score float64
	Human // Student继承Human
	// human Human // Student包含Human
}
