/*
*	定义二维数组，用于保存三个班级，每个班五名同学成绩。
*	并求出每个班级平均分、以及所有班级平均分
 */

package main

import (
	"fmt"
)

func main() {
	var arr [3][5]int
	var mark int
	var avg float64
	for i := 0; i < len(arr); i++ {
		mark1 := 0
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%v个班级，第%v名同学的成绩：\n", i+1, j+1)
			fmt.Scanln(&arr[i][j])
			mark += arr[i][j]
			mark1 += arr[i][j]
		}
		avg = float64(mark1) / 5
		fmt.Printf("第%v个班级的平均分为:%v\n", i+1, avg)
	}
	avg = float64(mark) / 15
	fmt.Printf("所有班级平均分为:%v\n", avg)
	fmt.Println(arr)

}
