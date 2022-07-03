package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person{
		Name:  "tom",
		Age:   19,
		Email: "asdas@qq.com",
	}
	str, _ := json.Marshal(person) // 结构体中的元素必须为大写
	fmt.Printf("json: %v\n", string(str))
	var p Person
	json.Unmarshal(str, &p) // 第二个参数为指针类型
	fmt.Printf("p: %v\n", p)
}
