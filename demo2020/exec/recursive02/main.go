/*
*	猴子每天吃桃堆里面的一半加一个，到第十天只有一个桃子了，求桃子总数
*/
package main

import(
	"fmt"
)

func peach(num int) int {
	if num == 10 {
		return 1
	}else{
		return (peach(num+1) + 1) * 2
	}
}

func main()  {
	fmt.Println("res=", peach(1))
}