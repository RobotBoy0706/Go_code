/*
* 有两个变量a和b，要求对其进行交换，但不能使用中间变量，然后输出
*/

package main

import(
	"fmt"
)

func main()  {
	var a int = 10
	var b int = 20

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%v b=%v", a, b)
}