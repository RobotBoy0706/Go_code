/*
*	随机生成10个随机数（1~100）保存到数组，并倒序打印以及求平均值，
*	求最大值和最大值的坐标，并查找是否存在55。
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randomSort [10]int
	var avg int
	var max int
	var maxIndex int
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixMicro())
		time.Sleep(1)
		randomSort[i] = rand.Intn(100)
		avg += randomSort[i]
	}

	fmt.Println(randomSort)
	fmt.Printf("该数组的平均值为%v \n", float64(avg/len(randomSort)))

	for i := 1; i < len(randomSort); i++ {

		fmt.Printf("该元素的倒序是%v \n", randomSort[10-i])
		if randomSort[i-1] > randomSort[i-2] {
			max = randomSort[i]
			maxIndex = i
		}
	}
	fmt.Printf("该数组的最大值为%v \n", max)
	fmt.Printf("该数组的最大值下标为%v \n", maxIndex)

}
