/*	参加百米运动会，如果用时8秒以内进入决赛，
*	否则提示淘汰。并且根据性别提示进入男子组或者女子组
*	输入成绩和性别进行判断
*/

package main

import (
	"fmt"
)

func main()  {
	var score int
	fmt.Println("请输入分数(秒)：")
	fmt.Scanln(&score)
	var sex int
	fmt.Println("请输入性别(1为男，0为女)：")
	fmt.Scanln(&sex)
	if score < 8 {
		if sex == 1 {
			fmt.Println("恭喜您进入了男子组决赛！")
		}else if sex == 0 {
			fmt.Println("恭喜您进入了女子组决赛！")
		}else{
			fmt.Println("不好意思，非法性别代号！")
		}
	}else{
		fmt.Println("很遗憾，您被淘汰了！")
	}
}