package main //封装可调用的函数 相当于 if __name__ ="__name__":

//导入包
import (
	"fmt"
)

func test1() {
	fmt.Println("Test2")
}

// if __name__ == "__name__"：
func main() {
	fmt.Println("Test1")
	test1()
}
