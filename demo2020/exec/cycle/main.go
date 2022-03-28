/*	成绩为100时，奖励一辆BMW；
*	成绩为（80，99]时，奖励一台iPhone
*	[60,80]奖励一个iPad
*	其他没有奖励
*	要求从键盘输入成绩并加以判断
*/

package main

import(
	"fmt"
)

func main()  {
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	if score == 100{
		fmt.Println("奖励一辆BMW！")
	}else if score > 80 && score <= 99 {
		fmt.Println("奖励一台iPhone！")
	}else if score >= 60 && score <= 80 {
		fmt.Println("奖励一个iPad")
	}else{
		fmt.Println("等着被揍！")
	}
}