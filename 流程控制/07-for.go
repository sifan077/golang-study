package main

import "fmt"

/*
* for 语句
 */

func main() {
	var count = 10
	fmt.Println("loop start")
	for count > 0 {
		fmt.Println(count)
		count--
	}
	fmt.Println("loop end")

	count = 10
	fmt.Println("loop start")
	// 无限循环
	for {
		fmt.Println(count)
		count--
		if count < 5 {
			break
		}
	}
	fmt.Println("loop end")

}
