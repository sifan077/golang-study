package main

import (
	"fmt"
	"strings"
)

/*
*  变量的声明与使用
 */

func main() {
	const PI = 3.14 // 声明一个常量，不可修改，修改就会报错
	// = 是赋值运算符，把右边的值赋给左边
	fmt.Println("圆周率", PI)
	var l = 5 // 声明一个变量可以修改
	l = 6
	var r = 6
	// 自加操作,没有 ++r,同样减也是这样
	r++
	r += 1
	r = r + 1
	fmt.Println("l=", r, " r=", r, " l*r=", l*r)

	var message = "Hi,shentu"
	var target = "shentu"
	var contains = strings.Contains(message, target)
	fmt.Println("target is in message ? ", contains)
}
