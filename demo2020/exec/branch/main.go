/*	
*	switch分支练习
*/
package main

import(
	"fmt"
)

func main()  {
	// 第一题:使用switch将小写a~e输出为大写，其他输出other
	// var alphabet byte
	// fmt.Println("请输入字母：")
	// fmt.Scanf("%c", &alphabet)
	// switch alphabet {
	// case 'a':
	// 	fmt.Println("A")
	// case 'b':
	// 	fmt.Println("B")
	// case 'c':
	// 	fmt.Println("C")
	// case 'd':
	// 	fmt.Println("D")
	// case 'e':
	// 	fmt.Println("E")
	// default:
	// 	fmt.Println("other")
	// }
	
	
	//	第二题: 输入的成绩大于60的输出“合格”，小于60输出“不合格”，输入不能大于100
	// var score int
	// fmt.Println("请输入成绩：")
	// fmt.Scanln(&score)
	// switch{
	// case score>=60 && score<=100:
	// 	fmt.Println("合格")
	// case score>0 && score<=60:
	// 	fmt.Println("不合格")
	// default:
	// 	fmt.Println("输入成绩不合法")
	// }

	// 第三题：根据用户指定月份，打印对应季节：3，4，5为春
	var month int
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("非法数字代号！")
	}

	// 第四题：
}