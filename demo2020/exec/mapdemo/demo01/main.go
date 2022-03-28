package main

import "fmt"

func main() {
	// 使用for-range的方式遍历
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no1"] = "天津"
	cities["no1"] = "上海"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v", k, v)
	}
}
