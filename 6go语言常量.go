package main

func main() {
	//常量的基础语法
	const name string = "张三"
	//常量的自动声明语法
	const name1 = "李四"
	//定义多常量赋值
	const name2, name3, name4 = "王五", "赵六", "苏七"
	//通过常量组实现枚举
	const (
		age   = 10
		sex   = "男"
		name6 = "huhs"
	)
	//通过常量组实现枚举(省略写法)
	const (
		num = 100
		num1
		num2
		num3
	)

}
