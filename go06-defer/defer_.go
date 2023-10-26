package main

import "fmt"

/**
defer 可以将某个函数入栈，等待当前函数生命周期结束再出栈执行
*/

func func1() {
	fmt.Println("func 1")
}

func func2() {
	fmt.Println("func 2")
}

func func3() {
	fmt.Println("func 3")
}

func deferFunc() {
	fmt.Println("defer func called...")
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	// return 语句会优先于defer函数执行
	defer deferFunc()
	return returnFunc()
}

func main() {
	// 以此入栈，在main函数执行结束之后依次出栈执行
	defer func1()
	defer func2()
	defer func3()

	fmt.Println("Start...")
	returnAndDefer()
}
