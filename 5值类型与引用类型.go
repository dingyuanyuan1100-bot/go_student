package main

import "fmt"

// 比如最容易混的切片和数组
// 数组是值类型,可以修改，无法增减。切片是引用类型，可以修改和增减。
func main() {
	var arrs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //切片写法
	var lit = [3]int{1, 2, 3}                   //数组写法
	lit[0] = 10                                 //数组修改
	lit1 := lit[:]                              //lit为数组，lit1为切片
	arrs = append(arrs, lit1...)                //切片扩容
	fmt.Println(arrs)
	fmt.Println(lit)
	lit2 := lit   //数组拷贝（赋值一份，互相独立）
	arrs1 := arrs //切片拷贝（底层共享，指针指向同一个内存）
	fmt.Println(arrs1)
	fmt.Println(lit2)

	//分类：
	//值类型：数组、int、struct、bool、float
	//引用类型：切片、map、channel
}
