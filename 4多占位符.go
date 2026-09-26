package main

import "fmt"

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	num    int
	bool_h bool
)

func main() {
	//var a int = 100           a := 100
	//var b float64 = 100.0		b := 100.0
	//var c string = "hello"	c := "hello"
	//var d string = "world"	d := "world"
	//var e bool = false		e := false

	//多变量声明短声明写法
	a, b, c, d, e := 100, 100.0, "hello", "world", false
	fmt.Printf(
		"int:%d\n"+ // %d为int整数占位
			"float64:%.1f\n"+ // %.1f为float64为浮点数占位
			"string:%s\n"+ // %s为字符串占位
			"value:%v\n"+ // %v为万能占位
			"bool:%t\n", // %t为布尔占位
		a, b, c, d, e,
	)
	fmt.Println(num, bool_h)
}
