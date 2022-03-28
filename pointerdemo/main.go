package main

import (
	"fmt"
)

func main() {

	// 基本数据类型的内存布局
	var i int = 10
	// 输出i的内存地址，&i
	fmt.Println("i的地址=", &i)

	// 下面的 var ptr *int = &i
	// 1、ptr是一个指针变量
	// 2、ptr的类型是 *int
	// 3、ptr本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v", &ptr)
	fmt.Printf("ptr 指向的值=%v", *ptr)
}
