package main

import "fmt"

// 本质是个指针
type AnimalI interface {
	Sleep()
	// 获取动物颜色
	GetColor() string
	// 获取动物种类
	GetType() string
}

type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalI) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor(), " type=", animal.GetType())
}

func main() {
	// 接口的数据类型，父类指针指向子类
	var animal AnimalI = &Cat{"Green"}
	showAnimal(animal)

	animal = &Dog{"Red"}
	showAnimal(animal)

	cat := Cat{"Yellow"}
	showAnimal(&cat)
	dog := Dog{"White"}
	showAnimal(&dog)
}
