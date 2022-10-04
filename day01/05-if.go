package main

import "fmt"

// if语句 和 或与非的使用

func main() {
	// go语言种  && || 同 java
	var year = 2100

	var flag = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if flag {
		fmt.Println(year, "年是闰年")
	} else {
		fmt.Println(year, "年不是闰年")
	}
}
