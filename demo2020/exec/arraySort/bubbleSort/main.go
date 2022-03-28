package main

import (
	"fmt"
)

// 冒泡排序
func BubbleSort(arrInt *[]int) []int {
	fmt.Println("排序前", *arrInt)
	temp := 0

	for i := 0; i < len(*arrInt)-1; i++ {

		for j := 0; j < len(*arrInt)-1-i; j++ {
			if (*arrInt)[j] > (*arrInt)[j+1] {
				temp = (*arrInt)[j]
				(*arrInt)[j] = (*arrInt)[j+1]
				(*arrInt)[j+1] = temp
			}
		}

	}
	fmt.Println("排序后", *arrInt)
	return *arrInt
}

func main() {
	var arrInt []int = []int{15, 65, 22, 1, 31}
	BubbleSort(&arrInt)
	fmt.Println("main arrInt:", arrInt)
}
