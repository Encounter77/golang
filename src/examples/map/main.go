package main

import (
	"fmt"
)

func main() {

	// 通过 make 创建 map1
	map1 := make(map[string]string, 10)
	// 设置 [k1:v1]
	map1["k1"] = "v1"
	fmt.Printf("map1 %+v\n", map1)
	// map2 value 为 function
	map2 := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	// 取 map2 的 value -> f 为 function
	f := map2["funcA"]
	// 得到 map2 的 value function 的 value
	fmt.Println("map2 value function call result", f())
	// 返回值 - 通过 exists 判断 key 是否存在
	value, exists := map1["k1"]
	if exists {
		fmt.Println(value)
	}
	// 常规遍历输出
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
