/*
*	创建一个byte类型的26个元素的数组，分别放置A~Z，使用for循环遍历打印。
*	提示：字符数据运算：'A' + 1 -> 'B'
 */

package main

import (
	"fmt"
)

func main() {
	var mychars [26]byte
	for i := 0; i < 26; i++ {
		mychars[i] = 'A' + byte(i)
	}

	for _, v := range mychars {
		fmt.Printf("%c \n", v)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", mychars[i])
	}
}
