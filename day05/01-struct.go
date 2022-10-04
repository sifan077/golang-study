package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

/**
* 1. 结构体的使用
* 2. struct 赋值是复制值传递
 */
type point struct {
	X float64
	Y float64
}

func distance(a, b point) float64 {
	t1 := (a.X - b.X) * (a.X - b.X)
	t2 := (a.Y - b.Y) * (a.Y - b.Y)
	return math.Sqrt(t1 + t2)
}

func main() {
	a := point{
		X: 3,
		Y: 4,
	}
	b := point{
		X: 0,
		Y: 0,
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(distance(a, b))

	// struct 转换为 json
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))

}
