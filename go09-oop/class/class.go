package main

import "fmt"

// 如果类名大写，表示其他包也可以访问。小写则不能
type Hero struct {
	// 如果类属性名大写，表示其他包也可以访问。小写则不能
	Name  string
	Ad    int
	level int
}

func (hero *Hero) Show() {
	fmt.Println("Name = ", hero.Name)
	fmt.Println("Ad = ", hero.Ad)
	fmt.Println("Level = ", hero.level)
}

func (hero *Hero) GetName() string {
	return hero.Name
}

// 如果方法名大写，表示其他包也可以访问。小写则不能
func (hero *Hero) setName(name string) {
	hero.Name = name
}

func main() {
	hero := Hero{
		Name:  "zhangsan",
		Ad:    100,
		level: 99,
	}

	fmt.Println("Hero Name: ", hero.GetName())

	hero.Show()
	hero.setName("lisi")
	hero.Show()
}
