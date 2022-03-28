/*
*	已知一个升序数列，插入一个元素，打印后依旧是升序
 */
package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{8, 12, 15, 20}
	var num int = 13

	for i := 0; i < len(arr); i++ {
		if num > arr[i] && num < arr[i+1] {
			for j := 0; j < len(arr)-i; j++ {
				arr[i], arr[i+1] = arr[i+1], arr[i]

			}

		}
	}
	fmt.Println(arr)
}
