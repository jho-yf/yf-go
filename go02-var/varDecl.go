package main

import "fmt"

/*
	四种变量的声明方式:
		方法一：声明一个变量 默认值为0
			var a int
		方法二：声明一个变量并初始化值
			var b string = "hello go"
		方法三：在初始化的时候，省略数据类型，通过值自动配置当前变量的数据类型
			var c = 100
		方法四：省去var关键字，直接自动匹配。常用，但只能在函数体内声明
			e := 3.14
*/

// 声明全局变量可以用方法一、二、三，使用方法四会编译报错
// 方法四只能在函数体内使用
var gA int
var gB string = "global"
var gC = 200

func main() {
	fmt.Println("gA = ", gA)
	fmt.Println("gB = ", gB)
	fmt.Println("gC = ", gC)

	// 方法一：声明一个变量 默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量并初始化值
	var b string = "hello go"
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// 方法三：在初始化的时候，省略数据类型，通过值自动配置当前变量的数据类型
	var c = 100
	fmt.Printf("type of c = %T\n", c)

	// 方法四：省去var关键字，直接自动匹配。常用
	d := "hello go"
	fmt.Printf("type of d = %T\n", d)
	e := 3.14
	fmt.Printf("type of e = %T\n", e)

	// 声明多个变量
	var aa, bb int = 100, 200
	fmt.Println("aa =", aa, ", bb =", bb)
	var id, name = 1, "jho"
	fmt.Println("id =", id, ", name =", name)

	// 多行多变量声明
	var (
		filename string = "a.txt"
		isExist  bool   = true
	)
	fmt.Println("filename =", filename, ", isExist =", isExist)
}
