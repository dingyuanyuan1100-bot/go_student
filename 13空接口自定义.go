package main

import "fmt"

func main() {
	var x interface{}                 //空接口自定义
	fmt.Printf("类型:%T, 值:%v\n", x, x) // 初始：<nil>

	x = 100
	fmt.Printf("类型:%T, 值:%v\n", x, x)

	x = "hello"
	fmt.Printf("类型:%T, 值:%v\n", x, x)

	x = 3.14
	fmt.Printf("类型:%T, 值:%v\n", x, x)
}
