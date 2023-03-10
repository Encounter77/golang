package main

import (
	"fmt"
)

// 切片是对数组一个连续片段的引用且支持动态扩展
// 数组定义中不指定长度即为切片
// 切片未初始化之前默认为nil，长度为0
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array", array)
	slice := array[1:3]
	fmt.Printf("part slice %+v\n", slice)
	fullSlice := array[:]
	fmt.Printf("full slice %+v\n", fullSlice)
	emptySlice := []int{}
	fmt.Printf("empty slice %+v\n", emptySlice)
	emptySlice = append(emptySlice, 9)
	emptySlice = append(emptySlice, 8)
	emptySlice = append(emptySlice, 7)
	emptySlice = append(emptySlice, 6)
	fmt.Printf("empty slice %+v\n", emptySlice)
	remove3rdSlice := deleteItem(emptySlice, 2)
	fmt.Printf("remove 3rd slice %+v\n", remove3rdSlice)

	// 构造切片 new & make
	// new
	// newSlice := new([]int)
	// make 可预设长度[2-length]和内存空间[3-capacity]，避免未来的内存拷贝
	// makeSlice := make([]int, 10, 20)

	// 修改切片的值
	slice1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("slice1", slice1)
	for _, value := range slice1 {
		// 值传递 - value 是切片元素的拷贝
		value *= 2
	}
	fmt.Println("modify slice1 by value", slice1)
	for index := range slice1 {
		// 下标传递(指针) - 修改的是传入切片的实际元素
		slice1[index] *= 2
	}
	fmt.Println("modify slice1 by index", slice1)
}

func deleteItem(slice []int, index int) []int {
	// 使用append将删除元素的前/后元素拼接实现元素删除
	return append(slice[:index], slice[index+1:]...)
}
