package main

import "fmt"

// 带有return的函数
func max_len(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//max_len为函数名，(num1, num2 int) 指函数输入num1 num2为int ,后面的int指的是返回为为int

// 多返回函数值
func swap(x, y string) (string, string) {
	return y, x
}

//(string, string)代表返回结果为两个都是字符串

func main() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}
