package array

import "fmt"

func ArrDemo() {

	// 固定长度的数组，默认值为0
	var arr1 [10]int

	fmt.Println("arr1数组:")
	for i := 0; i < 10; i++ {
		fmt.Print(arr1[i], " ")
	}
	fmt.Println()

	// 固定长度的数组并赋初始值
	arr2 := [5]int{1, 2, 3}
	arr2[4] = 1
	fmt.Println("arr2数组:")
	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i], " ")
	}
	fmt.Println()

	fmt.Println("for range遍历:")
	for index, value := range arr2 {
		fmt.Println("index=", index, " value=", value)
	}

	fmt.Println("查看数组的数据类型:")
	fmt.Printf("arr1 types: %T\n", arr1)
	fmt.Printf("arr2 types: %T\n", arr2)

}
