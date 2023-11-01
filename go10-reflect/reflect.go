package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", user)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is ", inputType)
	fmt.Printf("%T\n", inputType)

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is ", inputValue)

	// 通过type获取字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	user := User{1, "jho", 18}
	DoFiledAndMethod(user)
}
