package main

import (
	"fmt"
)

func just1() {
	coun22 := 100
	if coun22 <= 100 {
		fmt.Println("结果小于于等于100")
	}
	if coun22 > 100 {
		fmt.Println("大于100")
	}
}

// 传统if else写法
func just2() {
	coun23 := 100
	if coun23 <= 100 {
		fmt.Println("结果小于于等于100")
	} else if coun23 > 100 {
		fmt.Println("大于100")
	} else {
		fmt.Println("测试")
	}
}

// 表达式的 switch（值匹配），避免if else的冗余写法
func just3() {
	coun24 := 100
	switch coun24 {
	case 100:
		fmt.Println("结果为100")
	case 200:
		fmt.Println("结果为200")
	case 300:
		fmt.Println("结果为300")
	//多值判断
	case 400, 500, 600:
		fmt.Println("结果为400 || 500 || 600")
	}
}

// switch 的表达式返回 bool
func just4() {
	coun25 := 100
	switch coun25 >= 200 {
	case true:
		fmt.Println("结果大于等于200")
	case false:
		fmt.Println("结果小于200")
	}
}

// 无表达式 switch
func just5() {
	coun26 := 100
	switch {
	case coun26 <= 100:
		fmt.Println("结果小于等于100")
	case coun26 > 100:
		fmt.Println("结果大于100")
	}
}

// switch的fallthrough穿透写法
func just6() {
	coun27 := 100
	switch {
	case coun27 <= 100:
		fmt.Println("结果小于等于100")
		fallthrough //无视下一个判断条件，直接执行
	case coun27 > 100:
		fmt.Println("结果大于100")
	default:
		fmt.Println("结果不确定")
	}
}

// (x interface{})拆解
func just7(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x 是 int 类型")
	case string:
		fmt.Println("x 是 string 类型")
	case bool:
		fmt.Println("x 是 bool 类型")
	// 可以继续加任意多个 case
	default: // 可选分支
		fmt.Printf("未知类型: %T\n", x)
	}
}

func just8() {

}

func main() {
	just1()
}
