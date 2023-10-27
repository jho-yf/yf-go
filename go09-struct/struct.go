package main

import "fmt"

// 声明一种行的数据类型myint, 是int的一个别名
type myInt int

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func main() {
	var a myInt = 10
	fmt.Printf("a = %d\ntype of a = %T\n", a, a)

	var book Book
	book.title = "Golang"
	book.author = "zhangsan"
	fmt.Printf("%v\n", book)

	// 按值传递，无法修改结构体
	change1(book)
	fmt.Printf("%v\n", book)

	// 指针传递
	change2(&book)
	fmt.Printf("%v\n", book)
}

func change1(book Book) {
	// 按值传递
	book.author = "lisi"
}

func change2(book *Book) {
	// 指针传递
	book.author = "lisi"
}
