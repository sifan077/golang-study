package main

import "fmt"

/**
* slice 切片
 */

func main() {
	names := [...]string{
		"高子璇",
		"郝熙涵",
		"乔林",
		"罗秀兰",
		"王刚",
		"梁易轩",
		"顾泽惠",
		"宋润莎",
	}

	// 第一个切片, 切分数组的 [0-4)，左闭右开
	first := names[0:4]
	fmt.Println(first)
	// 第二个切片
	second := names[4:6]
	fmt.Println(second)
	// 第三个切片
	third := names[6:8]
	// 等价于  third := names[6:]
	fmt.Println(third)

	// 切分字符串
	str := "Hello"
	fmt.Println(str[:2])
	fmt.Println(str[3:])

	// 直接切分全部
	fmt.Println(str[:])

	// 声明一个切片
	arr := []int{1, 3, 4, 5, 7, 3, 74}
	// append添加
	arr = append(arr, 999)
	// 可以传入一个或者多个参数，第一个参数是估计的切片名
	arr = append(arr, 199, 299, 399)
	fmt.Println(arr)
	fmt.Println(len(arr), "是切片arr的长度")
	fmt.Println(cap(arr), "是切片arr的容量")

	// make函数
	// 创建arr1，容量是之前切片的两倍
	arr1 := make([]int, len(arr), (cap(arr))*2)
	fmt.Println(arr1)
	fmt.Println(len(arr1), "是切片arr1的长度")
	fmt.Println(cap(arr1), "是切片arr1的容量")

	// copy函数
	// 把arr的内容拷贝到arr1
	copy(arr1, arr)
	fmt.Println(arr1)
	fmt.Println(arr)
	fmt.Println(len(arr), "是切片arr的长度")
	fmt.Println(cap(arr), "是切片arr的容量")
}
