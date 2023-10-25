package main

import (
	"fmt"
	"yf-go/go04-init-and-import/lib1"        // 如果包被导入，但没有使用，编译会报错
	myLib2 "yf-go/go04-init-and-import/lib2" // 添加别名
	_ "yf-go/go04-init-and-import/lib3"      // 匿名导包：不被调用，编译时也不会报错，但会调用init方法
	. "yf-go/go04-init-and-import/lib4"      // 将fmt包中所有方法导入当前包中，调用的时候可以不加包名。不推荐使用
)

func main() {
	fmt.Println("main func...")
	lib1.Func()
	myLib2.Func()
	FuncOfLib4()
}

func init() {
	fmt.Println("main init...")
}
