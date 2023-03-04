package main

import (
	. "fmt" // 将fmt包中所有方法导入当前包中。不推荐使用
	"yf-go/go04-init-and-import/lib1"
	myLib2 "yf-go/go04-init-and-import/lib2" // 添加别名
	_ "yf-go/go04-init-and-import/lib3"      // 匿名导包：不被调用，编译时也不会报错，但会调用init方法
)

func main() {
	Println("main func...")
	lib1.Func()
	myLib2.Func()
}

func init() {
	Println("main init...")
}
