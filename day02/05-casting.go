package main

import (
	"fmt"
	"strconv"
)

/**
* 类型转换
 */

func main() {
	// 类型转换可能会丢失精度

	// 把数值转换为字符串
	var num = 45643452634
	fmt.Println(strconv.Itoa(num) + " Hello")
	fmt.Println(fmt.Sprintf("%v Hello", num))

	// 把字符串转换成bool
	res := "true"
	res1 := "yes"
	res2 := "1"
	fmt.Println(strconv.ParseBool(res))
	if res1 == "yes" {
		res1 = "t"
	}
	fmt.Println(strconv.ParseBool(res1))
	fmt.Println(strconv.ParseBool(res2))
	ans := "false"
	ans1 := "no"
	ans2 := "0"
	fmt.Println(strconv.ParseBool(ans))
	if res1 == "no" {
		res1 = "f"
	}
	fmt.Println(strconv.ParseBool(ans1))
	fmt.Println(strconv.ParseBool(ans2))

}
