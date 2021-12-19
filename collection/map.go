package main

import "fmt"

func main() {

	map1 := make(map[int]string)
	map1[123] = "44"
	for k, v := range map1 {
		fmt.Println("k:", k, "v:", v)
	}
	fmt.Println("map1[4]:", map1[4]) // 返回类型的默认值
	x, ok := map1[23]
	if ok {
		fmt.Println("map1[23]:", x)
	} else {
		fmt.Println("map1[23]不存在")
	}
	delete(map1, 100)
	delete(map1, 123)
	fmt.Println("map1:", map1)
	var map2 = map[int]string{111: "aaa", 222: "bbb"}
	fmt.Printf("map2: %v\n", map2)

}
