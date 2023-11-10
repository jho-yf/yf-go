package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func(a int, b int) bool {
		fmt.Println("a =", a, "b =", b)
		return true
	}(10, 20)

	go func() {
		// 用go创建承载一个形参为空，返回值为空的函数
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}

}
