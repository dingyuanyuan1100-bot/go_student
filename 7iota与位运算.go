package main

import "fmt"

// 基础iota用法
const (
	con = iota
	con1
	con2
	con3
	con4
)

// iota详细用法（独立值，iota+=1）
const (
	cont = iota
	cont1
	cont2 = "hi"
	cont3
	cont4 = "world"
	cont5 = "123"
	cont6
	cont7
)

// iota与位移运算
const (
	cont8  = 1 << iota //1<<0=1
	cont9  = 3 << iota //3<<1=6
	cont10             //3<<2=12
	cont11             //3<<3=24
)

func main() {
	fmt.Println(con, con1, con2, con3, con4)
	fmt.Println(cont, cont1, cont2, cont3, cont4, cont5, cont6)
}
