package main

import "fmt"

func main() {
	//指针写法
	var coun11 int = 100 //创建变量
	var per *int         //创建一个int指针变量
	per = &coun11        //&count11是获取其内存地址，让per指向count11的内存地址
	fmt.Println(per)     //打印出count11的内存地址

	//指针写法（简写法）
	count12 := 100
	per1 := &count12
	fmt.Println(per1)

	//指针分为 普通型指针 数组型指针 结构体型指针 函数指针 指针的指针（二级指针） nil指针(var pro *int)

}
