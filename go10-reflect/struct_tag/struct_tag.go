package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex" doc:"性别"`
}

func printTags(resume interface{}) {
	t := reflect.TypeOf(resume).Elem()

	for i := 0; i < t.NumField(); i++ {
		tags := t.Field(i).Tag
		fmt.Println(tags)

		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Println("info =", info, "doc =", doc)
	}
}

func main() {
	re := Resume{"jho", "male"}
	printTags(&re)
}
