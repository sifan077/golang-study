package main

import "fmt"

/**
* 函数的使用
 */

// Hello 无返回值的函数
func Hello(name string) {
	fmt.Println("Hello", name)
}

// Hi 返回值是string
func Hi(name string) string {
	var res = "Hi," + name
	return res
}

// Hint 多返回值的函数
func Hint(name string) (string, int) {
	return name, len(name)
}

func main() {
	Hello("shentu")
	fmt.Println(Hi("shentu"))

	var name, length = Hint("shentu")
	fmt.Println(name, length)
}
