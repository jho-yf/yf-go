package main

import (
	"fmt"
	"yf-go/go07-slice/array"
	"yf-go/go07-slice/slice"
	"yf-go/go07-slice/slice_op"
)

func main() {
	fmt.Println("========================Array Demo========================")
	array.ArrDemo()
	fmt.Println("========================Slice Demo========================")
	slice.SliceDemo()
	fmt.Println("========================Slice Op Demo========================")
	slice_op.SliceOpDemo()
}
