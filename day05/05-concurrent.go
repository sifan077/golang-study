package main

import (
	"fmt"
	"time"
)

/**
*  go语言携程的使用
 */

func printStr(msg string){
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(1000)
	}
}

func main() {
	// 开启并发，多线程执行
	go printStr("Hello")
	printStr("World")
}
