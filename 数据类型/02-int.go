package main

import "fmt"

/**
* 整数的使用
 */
func main() {
	// 声明整数变量
	var a = 15

	var b int = 16

	fmt.Println(a, " ", b)

	// 打印数据类型
	fmt.Printf("%T\n", a)

	// 声明16进制的数
	var c = 0x8d
	fmt.Println("0x8d=", c)

	// 整数范围会环绕

	// Math包里声明了，int最大和最小值

}
