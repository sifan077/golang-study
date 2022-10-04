package main

import "fmt"

/**
*  map的使用，map是引用传递，修改就会修改原来的
 */

func main() {
	// 声明map     var hash = map[type_of_key] value_of_key
	var hash = map[string]int{
		"shentu": 1,
		"sifan":  2,
	}
	fmt.Println(hash["shentu"])
	// 插入key-value值
	hash["lan"] = 3

	// 检查map中是否存在key
	// map中不存在的key值，value是 默认零值
	// 获取值会返回两个值，一个是value,一个是bool值，告诉是否存在于map中
	val, ok := hash["lan"]
	if ok {
		fmt.Println("lan对应的value:", val)
	} else {
		fmt.Println("hash中不存在 ", val)
	}

	// 删除key值
	delete(hash, "lan")
	// 测试删除
	val, ok = hash["lan"]
	if ok {
		fmt.Println("lan对应的value:", val)
	} else {
		fmt.Println("hash中不存在 ", val)
	}
}
