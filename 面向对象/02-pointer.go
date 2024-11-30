package main

import "fmt"

/**
*  指针
 */

func main() {
	var a int = 20
	var p *int
	p = &a
	fmt.Println("a的地址是", &a)
	fmt.Println("p指针储存的位置是", p)
	fmt.Println("p变量的值是", *p)
	// 声明一个空指针
	var ptr *int
	if ptr == nil {
		fmt.Println("ptr 是空指针")
	}
	if ptr != nil {
		fmt.Println("ptr 不是空指针")
	}

	var b int = 100
	var c int = 200

	fmt.Printf("交换前 a 的值 : %d\n", b)
	fmt.Printf("交换前 b 的值 : %d\n", c)

	/* 调用函数用于交换值
	 * &b 指向 b 变量的地址
	 * &c 指向 c 变量的地址
	 */
	swap(&b, &c)

	fmt.Printf("交换后 a 的值 : %d\n", b)
	fmt.Printf("交换后 b 的值 : %d\n", c)
}

func swap(x *int, y *int) {
	// 可以写成 *x, *y = *y, *x
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
