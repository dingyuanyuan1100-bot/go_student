package main //封装可调用的函数 相当于 if __name__ ="__name__":

import (
	"fmt"
)

func test1() {
	fmt.Println("Test2")
}
func main() {
	fmt.Println("Test1")
	test1()
}
