package main

import "fmt"

/**
* 函数的传递和使用
 */

func hi() string {
	return "Hi"
}

// 函数中传入函数

func hello(name string, out func(nick string)) {
	out(name)
}

func sout(name string) {
	fmt.Println(name)
}

var f3 = func() {
	fmt.Println("这是一个匿名函数")
}

func main() {
	// 把函数赋值给变量
	var say = hi
	// 把变量函数执行了
	fmt.Println(say())

	hello("shentu", sout)
	// 函数的调用
	f3()

	var fu3 = func() {
		fmt.Println("main函数里的匿名函数")
	}
	fu3()

	// 直接执行的匿名函数
	func(){
		fmt.Println("直接执行的匿名函数")
	}()

}
