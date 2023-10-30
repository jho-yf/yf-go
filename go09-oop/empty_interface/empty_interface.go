package main

import "fmt"

// interface{} 万能数据类型
func myFunc(arg interface{}) {

	// interface{} 断言
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string, value=", value)
	} else {
		fmt.Printf("arg is %T", arg)
		fmt.Println(", value=", arg)
	}
}

type Book struct {
	author string
}

func main() {
	book := Book{"zhangsan"}
	myFunc(book)
	myFunc("test")
	myFunc(555)
	myFunc(3.14)
}
