package main

import "fmt"

func main() {
	var coun uint = 10
	var coun1 uint = 20

	var coun2 = coun & coun1
	fmt.Println(coun2)
}

//注意位运算
//	10110
// &10001
//_________
//  10000

//注意位运算
//	10110
// |10001
//_________
//  10111
