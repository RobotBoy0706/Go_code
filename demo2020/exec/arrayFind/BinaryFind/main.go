package main

import (
	"fmt"
)

func BinaryFind(arr *[7]int, leftIndex int, rightIndex int, valFind int) {

	if leftIndex > rightIndex {
		fmt.Println("未找到该元素！")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > valFind {
		BinaryFind(arr, leftIndex, middle-1, valFind)
	} else if (*arr)[middle] < valFind {
		BinaryFind(arr, middle+1, rightIndex, valFind)
	} else {
		fmt.Printf("该元素下标为%v \n", middle)
	}
}

func main() {
	arr := [7]int{5, 12, 21, 29, 35, 39, 85}
	BinaryFind(&arr, 0, len(arr)-1, 29)

}
