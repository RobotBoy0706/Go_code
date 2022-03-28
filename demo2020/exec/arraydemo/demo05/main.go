package main

import (
	"fmt"
)

func main() {
	// 用append内置函数，可以对切片进行动态追加
	var slice []int = []int{100, 200, 300}
	// 通过append直接给slice追加具体的元素
	slice = append(slice, 400, 500)
	// 通过append将切片slice追加给slice
	slice = append(slice, slice...)
	fmt.Println("slice", slice)

	// 切片拷贝
	var slice2 = make([]int, 15)
	copy(slice2, slice)
	fmt.Println("slice2", slice2)
}
