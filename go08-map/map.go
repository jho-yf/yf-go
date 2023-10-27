package main

import "fmt"

func main() {
	// 声明map
	// 方式一 先声明，后使用make分配数据空间
	var map1 map[string]string
	fmt.Println("map1是否为一个空map:", map1 == nil)
	map1 = make(map[string]string, 10)
	fmt.Println("map1是否为一个空map:", map1 == nil)

	map1["one"] = "java"
	map1["two"] = "go"
	map1["three"] = "python"
	fmt.Println(map1)

	// 方式二 声明并分配空间
	map2 := make(map[int]string)
	map2[0] = "java"
	map2[1] = "go"
	map2[2] = "python"
	fmt.Println(map2)

	// 方式三 声明并赋值
	map3 := map[int]string{
		0: "java",
		1: "go",
		2: "python",
	}
	fmt.Println(map3)

	// map基本使用方式
	cities := make(map[string]string)

	// 添加
	cities["China"] = "Beijing"
	cities["Japan"] = "Tokyo"
	cities["USA"] = "NewYork"

	// 修改
	cities["USA"] = "DC"

	// 删除
	delete(cities, "USA")

	printMap(cities)
}

func printMap(cities map[string]string) {
	// 按值传递
	// 遍历
	for key, value := range cities {
		fmt.Println("key=", key, " value=", value)
	}
}
