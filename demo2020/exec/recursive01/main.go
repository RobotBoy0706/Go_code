/*
*	用递归求出斐波那契数：1, 1, 2, 3, 5, 8, 13....
*/
package main

import(
	"fmt"
)

func test(n int) int {
	if n == 1 || n ==2 {
		return 1
	}else{
		return test(n-1) + test(n-2)
	}
}

func main()  {
	
	fmt.Println("res=", test(14))
	
	fmt.Println("res=", test1(5))
}

func test1(num int) int {
	if num == 1 {
		return 3
	}else{
		return 2 * test1(num-1) +1
	}
}