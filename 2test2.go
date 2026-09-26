package main

import (
	"fmt"
)

// var是变量声明，可以改变其值
var x float64 = 42

// const是常量值，不可以改变其值
const y float64 = 42.0

// 变量与常量测试，字符串拼接测试
func test2() {
	fmt.Println("Test")
	fmt.Println("hi" + "➕➕1")
	x += 1
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + y)
}
func test3(a float64, b float64) {
	c := 10.0 //短声明只能声明变量，不能声明常量(但是c:=10默认为int,c:=10.0默认为float64。float64占用8字节，float32占用4字节)
	if a == b {
		fmt.Println("这是测试3")
		test2()
		fmt.Println("Test")
		fmt.Println(a + b + c)
	}

}

func main() {
	test3(x, y)
}
