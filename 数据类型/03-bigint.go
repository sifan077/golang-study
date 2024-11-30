package main

import (
	"fmt"
	"math/big"
)

/**
* 大数处理
 */

func main() {
	//int64  < float64
	// 指数形式的默认类型是float64
	// 超过10^18次方使用big包
	var distance = new(big.Int)
	distance.SetString("246464618694134654123456", 10)
	var speed = new(big.Int)
	speed.SetString("30010024124021", 10)
	var second = new(big.Int)
	second.Div(distance, speed) // 相当于 second = distance/speed
	fmt.Println(second)

	var res = big.NewInt(86400)
	fmt.Println(res)

	var ans = new(big.Int)
	ans.SetString("86400", 10)
	fmt.Println(ans)
}
