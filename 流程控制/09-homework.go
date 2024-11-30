package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 添加随机种子，实现真随机
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		switch a := rand.Intn(3) + 1; a {
		case 1:
			fmt.Print("Space ")
		case 2:
			fmt.Print("SpaceX ")
		case 3:
			fmt.Print("Virgin  ")
		default:
			fmt.Print("void  ")
		}
		fmt.Print(rand.Intn(30)+1, "  ")
		if i := rand.Intn(2) + 1; i == 1 {
			fmt.Print("One-way  ")
		} else {
			fmt.Print("Round-trip  ")
		}
		fmt.Println(rand.Intn(1400) + 3601)
	}
}
