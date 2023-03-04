package lib1

import "fmt"

// 函数名首字母大写表示函数能够提供给外部调用
// 函数名首字母小写表示函数是私有的，只能给内部调用
func Func() {
	fmt.Println("lib1 func...")
}

func init() {
	fmt.Println("lib1 init...")
}
