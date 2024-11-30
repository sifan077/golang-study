package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*  短声明声明变量, 变量的作用范围
 */

// 全局变量
var message = "Hello"

func main() {
	fmt.Println("package级别的变量:", message)
	qq := "短声明变量只可以在函数内声明"
	fmt.Println(qq)
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// 添加随机种子，实现真随机
	rand.Seed(time.Now().UnixNano())
	if i := rand.Intn(100) + 1; i%2 == 0 {

		fmt.Println(i, "是偶数")
	} else {
		fmt.Println(i, "不是偶数")
	}
}
