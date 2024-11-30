package main

import "fmt"

/**
 * 实数
 */

func main() {
	// 声明浮点数变量
	var a = 3.1415926

	// 如果声明的数是一个整数，必须声明 float64才能变为浮点数,否则就是 int 变量
	// float64 = double 双精度 8字节   float32 单精度 4字节
	var b float64 = 3

	fmt.Println(a, " ", b)

	// go 语言每个类型都有默认值，称作零值
	// 声明变量却不初始化，它的值就是零值

}
