package slice

import "fmt"

func SliceDemo() {
	// 切片slice的声明方式
	// 方式一：初始化并指定值
	slice1 := []int{1, 2, 3}
	fmt.Printf("slice1 = %v, len = %d\n", slice1, len(slice1))
	fmt.Println("slice1是否为空: ", slice1 == nil)

	// 方式二：使用make函数分配空间，初始化值为空字符
	var slice2 []string = make([]string, 5)
	fmt.Printf("slice2 = %v, len = %d\n", slice2, len(slice2))
	fmt.Println("slice2是否为空: ", slice2 == nil)

	// 方式三：利用类型推导
	slice3 := make([]int, 5)
	fmt.Printf("slice3 = %v, len = %d\n", slice3, len(slice3))
	fmt.Println("slice3是否为空: ", slice3 == nil)

	// 方式四：只声明不分配空间
	var slice4 []int
	fmt.Printf("slice4 = %v, len = %d\n", slice4, len(slice4))
	fmt.Println("slice4是否为空: ", slice4 == nil)

	// 数组
	arr1 := [10]int{1, 2, 3}
	fmt.Printf("arr1 type is %T\n", arr1)
	printArr(arr1)

	// 动态数组，切片
	slice5 := []int{1, 2, 3, 4}
	fmt.Printf("slice4 type is %T\n", slice5)
	printSlice(slice5)
}

func printArr(arr [10]int) {
	// 按值传递
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func printSlice(slice []int) {
	// 引用传递
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
