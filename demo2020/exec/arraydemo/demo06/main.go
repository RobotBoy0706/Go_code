/*
*	编写一个函数 fbn(n int) 能够接收斐波那契数列并放入切片中。
*	形式：arr[0] = 1; arr[1] = 1; arr[2] = 2; arr[3] = 3; arr[4] = 5
 */

package main

import (
	"fmt"
)

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-2] + fbnSlice[i-1]
	}
	return fbnSlice
}

func main() {
	fbnSlice := fbn(10)
	fmt.Println(fbnSlice)
}
