package main

import "fmt"

/*
* switch语句的使用
 */

func main() {
	var score = 'C'
	switch score {
	case 'A':
		fmt.Println("优秀")
	case 'B':
		fmt.Println("良好")
	case 'C':
		fmt.Println("中等")
		fallthrough //加上 fallthrough 会自动执行下一个case
	case 'D':
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
