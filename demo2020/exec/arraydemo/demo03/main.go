package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [5]int{12, 13, 65, 75, 52}
	maxVal := intArr[0]
	maxValIndex := 0
	sum := 0
	for i, v := range intArr {
		// 求最大值
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}

		// 累计求和
		sum += v
	}
	fmt.Printf("maxVal=%v maxValIndex=%v", maxVal, maxValIndex)
	fmt.Printf("sum=%v 平均值=%v", sum, float64(sum)/float64(len(intArr)))
}
