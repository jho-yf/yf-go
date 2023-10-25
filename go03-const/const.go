package main

import "fmt"

// const() 可以用来定义枚举
const (
	BEIJING   = 0
	SHANGHAI  = 1
	GUANGZHOU = 2
	SHENZHEN  = 3
)

// 可以在const() 添加 `iota` 关键字，每行的iota都会累加1，第一行的iota默认值是0
const (
	SECOND = iota
	MINUTE
	HOUR
)

const (
	RUNNING = iota * 10
	RUNNABLE
	TERMINAL
)

const (
	a, b = iota + 1, iota + 2
	c, d = iota * 2, iota * 3
	e, f = iota, iota
)

func main() {
	// 常量：只读属性
	const len int = 10
	fmt.Println("len =", len)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("GUANGZHOU = ", GUANGZHOU)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("====================")
	fmt.Println("SECOND = ", SECOND)
	fmt.Println("MINUTE = ", MINUTE)
	fmt.Println("HOUR = ", HOUR)

	fmt.Println("====================")
	fmt.Println("RUNNING = ", RUNNING)
	fmt.Println("RUNNABLE = ", RUNNABLE)
	fmt.Println("TERMINAL = ", TERMINAL)

	fmt.Println("====================")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
}
