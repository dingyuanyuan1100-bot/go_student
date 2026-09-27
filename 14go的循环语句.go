package main

import "fmt"

//Go for 循环三部分执行顺序
//for 初始化; 条件; 循环后表达式 {
//循环体
//}

func just15() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

//just15 与 just16（ 等价写法 ）

func just16() {
	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}
	//简化写法
	i := 1
	for i < 5 {
		i++
	}
	fmt.Println("成功")
}

// goto语句的使用
func just17() {
	var a int = 10
loop:
	for a < 20 {
		a += 1
		goto loop
	}

}

func main1() {

}
