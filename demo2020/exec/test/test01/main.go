/*
*	判断一个数是否为水仙花数（一个三位数，其各个位上的数字立方和等于其本身）
*/

package main

import(
	"fmt"
	"math"
)

func main()  {
	fmt.Println("请输入需要判断的数字（三位数）：")
	var num int
	fmt.Scanln(&num)
	// var num1 int = num / 100
	// var num2 int = (num - num1 * 100) / 10
	// var num3 int = num - num1 * 100 - num2 * 10
	// var formula int
	// if num%100 >= 1 && num%100 <10 {
	// 	for i := 0; i < 3; i++ {
			
	// 	}
	// }else{
	// 	fmt.Println("输入的数字不在范围内！")
	// }
	str := string(num)
	var num2 float64
	if num%100 >= 1 && num%100 <10 {
		for i := 0; i < len(str); i++ {
			// j := int(str[i])
			num2+=math.Pow(float64(str[i]), 3)
		}
	}else{
	 	fmt.Println("输入的数字不在范围内！")
	 }
	
}

