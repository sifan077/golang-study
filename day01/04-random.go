package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*  猜数游戏, 生成随机数
 */

func main() {
	// 添加随机种子，实现真随机
	rand.Seed(time.Now().UnixNano())
	var num = rand.Intn(100) + 1
	fmt.Println("随机生成的数是", num)
	var guess = 0
	for {
		fmt.Print("请输入您猜测的值:")
		fmt.Scan(&guess)
		if guess == num {
			fmt.Println("您猜对了")
			break
		} else if guess > num {
			fmt.Println("您猜的数大了")
		} else {
			fmt.Println("您猜的数小了")
		}
	}
}
