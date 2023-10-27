package slice_op

import "fmt"

func SliceOpDemo() {
	// 定义一个切片，长度为3，容量为5
	var numbers = make([]int, 3, 5)
	fmt.Printf("%v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))

	// 添加一个元素
	numbers = append(numbers, 3)
	fmt.Printf("%v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))

	numbers[0] = 0
	numbers[1] = 1
	numbers[2] = 2
	// 切片截取
	numbers = numbers[:]
	fmt.Printf("%v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))
	numbers = numbers[1:4]
	fmt.Printf("%v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))
	numbers = numbers[1:]
	fmt.Printf("%v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))
	numbers = numbers[:1]
	fmt.Printf("%v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))

	// 拷贝
	var tmp []int = make([]int, 1)
	copy(tmp, numbers)
	fmt.Printf("%v, len=%d, cap=%d\n", tmp, len(tmp), cap(tmp))
}
