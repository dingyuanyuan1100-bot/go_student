package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// 遍历切片
func main() {
	//数组或切片便利
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	//只遍历value
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	//只遍历key
	for key := range pow {
		fmt.Printf("%d\n", pow[key])
	}
}
