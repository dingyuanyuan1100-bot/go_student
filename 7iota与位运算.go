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

func main() {
	fmt.Println(con, con1, con2, con3, con4)
	fmt.Println(cont, cont1, cont2, cont3, cont4, cont5, cont6)
}
