package main

import "fmt"

/*
* 多语言文本
 */
func main() {
	var message = "Hello World"
	//字符串不可直接使用数组的形式获取更改
	fmt.Println(message)

	fmt.Println("message的长度为", len(message))

	fmt.Println("正常的换行符，\n输出的时候不会显示")

	fmt.Println(`不正常的换行符\n,会输出出来，因为转义了`)

	var c = message[3]
	fmt.Printf("%c\n", c)

	// range遍历字符串
	for i, c := range message {
		fmt.Printf("%d %c\n", i, c)
	}

}
