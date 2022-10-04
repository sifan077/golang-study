package main

import (
	"errors"
	"fmt"
	"math"
)

/**
*  异常处理
 */

type error interface {
	Error() string
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}

	return math.Sqrt(f), nil
}
func main() {
	result, err := Sqrt(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
