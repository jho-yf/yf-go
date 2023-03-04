package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// 多个返回值，匿名
func swap(a string, b string) (string, string) {
	return b, a
}

// 多个返回值，带形参名称
func incr(a int, b int) (num1 int, num2 int) {
	num1 = a + 1
	num2 = b + 1
	return
}

// 返回值形参有默认值
func incr1(a int, b int) (num1, num2 int) {
	return
}

func main() {
	sum := add(1, 1)
	fmt.Println(sum)
	fmt.Println(add(2, 2))

	a, b := swap("world", "hello")
	fmt.Println(a, b)

	n1, n2 := incr(99, 999)
	fmt.Println(n1, n2)

	n1, n2 = incr1(9, 9999)
	fmt.Println(n1, n2)
}
