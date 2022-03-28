package main

import "fmt"

func main() {

	// 1、 声明一个动态数组
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	// 2、增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "孙猴子"
		monsters[1]["age"] = "1000"
	}
	newMonster := map[string]string{
		"name": "新的妖怪~蜘蛛精",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

}
