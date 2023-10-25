package main

import "fmt"

func changeValue(p *int) {
	*p = 10
}

func swap(pNum1 *int, pNum2 *int) {
	var temp int = *pNum1
	*pNum1 = *pNum2
	*pNum2 = temp
}

func main() {
	var a int = 1
	changeValue(&a)
	fmt.Println("a = ", a)

	// swap
	var num1 int = 10
	var num2 int = 20

	fmt.Println("num1 = ", num1, "num2 = ", num2)

	swap(&num1, &num2)

	fmt.Println("num1 = ", num1, "num2 = ", num2)

	// 定义指针
	var p *int = &num1
	// 打印指针地址
	fmt.Println(&num1)
	fmt.Println(p)
	fmt.Println(p == &num1)

	// 二级指针
	var pp **int = &p
	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println(&p == pp)

	// 三级指针
	fmt.Println(&pp)
}
