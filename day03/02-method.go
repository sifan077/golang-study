package main

import "fmt"

/**
* 方法的使用
 */

// celsius  摄氏度
type celsius float64

// kelvin 开尔文温度
type kelvin float64

type cc float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// 通过方法添加行为
// 声明一个kelvin变量 var k kelvin = 15
// k.toCelsius(k)调用方法的样例
func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}


func main() {
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)
	fmt.Println(kelvin.toCelsius(15))
	var k kelvin = 15
	fmt.Println(k.toCelsius())
	fmt.Println(kelvinToCelsius(15))
}
