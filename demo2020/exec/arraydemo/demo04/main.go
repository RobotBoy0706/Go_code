/*
*	随机生成五个数，并将其反转打印
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr [5]int
	len := len(intArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr[i] = rand.Intn(100)
	}

	fmt.Println("交换前=", intArr)

	// 反转打印，交换的次数是 len/2
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr[len-1-i]
		intArr[len-1-i] = intArr[i]
		intArr[i] = temp
	}

	fmt.Println("交换后=", intArr)
}
