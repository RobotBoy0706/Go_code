package main

import (
	"fmt"
)

/*
*	一个养鸡场有六只鸡，体重分别是3, 5, 1, 3.4, 2, 50kg。求总体重和平均重
 */

func main() {
	// 1、定义一个数组
	var hens [6]float64
	// 2、给数组每一个元素赋值，元素的下标是从0开始的
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	// 3、遍历数组求总体重
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	// 4、求出平均体重
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v avgWeight=%v", totalWeight, avgWeight)
}
