/*
*	1、假如还有97天放假，问：xx个星期零几天
*	2、定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：
*	5/9*（华氏温度-100），请求出华氏温度对应的摄氏温度。
*/

package main

import(
	"fmt"
)

func main()  {
	var date int = 97
	var weeks int = date / 7
	var days int = date % 7
	fmt.Println("还有",weeks,"个星期零",days,"天")

	var hs_temp int = 134
	var ss_temp float32 = 5.0 / 9 * (float32)(hs_temp - 100)
	fmt.Println("华氏温度",hs_temp,"对应的摄氏温度是",ss_temp)
}