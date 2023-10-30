package main

import "fmt"

type Human struct {
	name string
	sex  bool
}

func (human *Human) GetName() string {
	return human.name
}

func (human *Human) GetSex() bool {
	return human.sex
}

type SuperMan struct {
	// 继承Human
	Human

	level int
}

func (superMan *SuperMan) GetLevel() int {
	return superMan.level
}

// 覆盖父类方法
func (superMan *SuperMan) GetName() string {
	return "S " + superMan.name
}

func main() {
	superman := SuperMan{
		Human{name: "zhangsan", sex: true},
		99,
	}
	fmt.Println(superman.GetName())
	fmt.Println(superman.GetSex())
	fmt.Println(superman.GetLevel())

	var superman2 SuperMan
	superman2.name = "lisi"
	superman2.sex = false
	superman2.level = 100
	fmt.Println(superman2.GetName())
	fmt.Println(superman2.GetSex())
	fmt.Println(superman2.GetLevel())
}
