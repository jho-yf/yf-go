package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	// 序列化
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("%s\n", jsonStr)

	// 反序列化
	movie1 := Movie{}
	err = json.Unmarshal(jsonStr, &movie1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", movie1)

}
