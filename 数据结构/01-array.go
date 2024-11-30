package main

import "fmt"

/**
*  数组的声明与使用
 */

func main() {
	// 未赋值的默认是一个空字符串
	// 数组传入函数，是值传递，不是地址传递，修改不会影响原来的数组
	var names [8]string
	names[0] = "shentu"
	names[1] = "sifan"
	names[2] = "lan"

	t := names[2]
	fmt.Println(t)

	fmt.Printf("%v\n", len(names))

	fmt.Println(names[3] == "")

	// 直接声明一个数组
	arr := [5]int{1, 3, 5, 8, 20}
	// 可以写成 arr := [...]int{1, 3, 5, 8, 20}
	fmt.Println(arr)

	// for循环遍历
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
	// range循环遍历
	for i, item := range arr {
		fmt.Println(i, item)
	}

	// 二维数组声明
	var board [8][8]string
	board[0][0] = "h"
	board[1][1] = "h"
	board[2][2] = "h"
	board[3][3] = "h"
	board[7][7] = "r"
	for i, item := range board {
		fmt.Printf("%v %v", i, item)
	}

}
