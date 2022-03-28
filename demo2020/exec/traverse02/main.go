/*
*	打印1~100所有是9的倍数的数以及求和
*/

package main

import(
	"fmt"
)

func main()  {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		if i%9==0 {
			fmt.Println(i)
			sum+=i
		}
	}
	fmt.Println("符合条件数总和为：", sum)

	var n int = 6
	for i := 0; i < n; i++{
		fmt.Printf("%v + %v = %v \n", i, n-i, n)
	}
}		