package main

// 多变量声明
var a, b int = 1, 2
var c, d = 3, 4

// 声明一个变量默认为0
var s string  //等于 s==""
var e int     //等于 e==0
var f bool    //等于 f==false
var g float64 //等于 g==0

func main() {
	println(a, b)
	println(c, d)
	println(s)
	println(e)
	println(f)
	println(g)
}
