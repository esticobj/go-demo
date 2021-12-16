package main

import "fmt"

func main() {
	// 创建定长数组
	arr := [5]int{1, 2, 3}
	fmt.Println("使用fori遍历")
	for i := 0; i < len(arr); i++ {
		fmt.Println("i:", i, ",j:", arr[i])
	}
	fmt.Println("使用range遍历")
	for k, v := range arr {
		v += 1 // 修改v不会修改数组的值，v只是每次遍历元素的副本，并且只会申请一次存储空间
		fmt.Println("k:", k, ",v:", v, ",addr:", &v)
	}
	fmt.Println(arr)
	// 创建切片（不定长数组）
	names := []string{"a", "b", "c"} //[]里不指定长度就是切片
	fmt.Println("names len:", len(names), "cap:", cap(names))
	names1 := append(names, "d") //names 不会变化，names1容量不够会会扩容到原来capcity的2倍
	fmt.Println(names)
	fmt.Println(names1)
	fmt.Println("names len:", len(names), "cap:", cap(names))
	fmt.Println("names1 len:", len(names1), "cap:", cap(names1))
	for i := 0; i < 50; i++ {
		//names := append(names, "d") //内部变量会覆盖外部变量
		names = append(names, "d") //names容量不够会扩容到原来capcity的2倍
		fmt.Println("names len:", len(names), "cap:", cap(names))
	}

	//基于定长数组创建切片
	x := [3]int{1, 2, 3}
	x1 := x[0:3]
	fmt.Println("x1:", x1)
	x1[0] = 5
	fmt.Println("x:", x) //x也会跟着修改，x1是x的浅拷贝
	fmt.Println("x1:", x1)

	//创建带有初始容量的切片
	ls := make([]int, 7) // len=cap=7
	fmt.Println("ls: len:", len(ls), "cap:", cap(ls))
	//创建带有长度和初始容量的切片
	ls2 := make([]int, 0, 7) // len=0,cap=7
	fmt.Println("ls: len:", len(ls2), "cap:", cap(ls2))

	x2 := make([]int, len(x)) // len为0会报错
	//深拷贝
	copy(x2, x[:]) //第二个参数必须是切片
	fmt.Println("x2:", x2)
	x2[0] = 9
	fmt.Println("x2:", x2)
	fmt.Println("x1:", x1)

}
