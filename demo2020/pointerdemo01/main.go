/*
*	写一个程序：获取一个int变量num的地址，并显示到终端。
*	将num的地址赋给指针ptr，并通过ptr去修改num的值
*/

package main

import (
	"fmt"
)

func main()  {
	var num int = 1
	fmt.Printf("当前num的地址为%v", &num)
	
	var ptr *int = &num
	*ptr = 2
	fmt.Printf("当前num的地址为%v", num)
}
	
