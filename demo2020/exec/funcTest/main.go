/*
*	编写函数，可以计算两个数的和、差并返回结果
*/

package main

import(
	"fmt"
)

// 函数，用来计算并返回结果
func CountFunc(num1 int, num2 int) (int, int) {
	// var count int
	// if operator == '+' {
	// 	count = num1 + num2
	// 	return count
	// }else if operator == '-' {
	// 	count = num1 - num2
	// 	return count
	// }else{
	// 	return fmt.Println("暂不支持该种运算！")
	// }
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

func main()  {
	res1, res2 := CountFunc(1, 2)
	fmt.Printf("res1=%v res2=%v\n", res1, res2)
}