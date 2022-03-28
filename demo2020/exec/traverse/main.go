/*
*	使用for循环遍历字符串并输出
*/

package main

import(
	"fmt"
)

func main()  {
	var str string = "HelloWorld,北京!"
	for i := 0; i < len(str); i++{
		fmt.Printf("%c \n", str[i])
	}

	str2 := []rune(str)
	for i := 0; i < len(str2); i++{
		fmt.Printf("%c \n", str2[i])
	}

	str = "ILoveChina北京"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}