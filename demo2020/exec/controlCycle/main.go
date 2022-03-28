/*
*	a~g代表星期一到星期天，用控制循环表示出来
*/

package main

import(
	"fmt"
)

func main()  {
	var data1 byte
	fmt.Println("请输入字母（仅限a~g）：")
	fmt.Scanf("%c", &data1)
	switch data1 {
	case 'a' :
		fmt.Println("今天是星期一。")
	case 'b' :
		fmt.Println("今天是星期二。")
	case 'c' :
		fmt.Println("今天是星期三。")
	case 'd' :
		fmt.Println("今天是星期四。")
	case 'e' :
		fmt.Println("今天是星期五。")
	case 'f' :
		fmt.Println("今天是星期六。")
	case 'g' :
		fmt.Println("今天是星期七。")
	default :
		fmt.Println("输入的代号非法！")
	}
}