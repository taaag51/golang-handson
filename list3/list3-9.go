package main

import (
	"fmt"
)

// MyData is structure.
type MyData struct {
	Name string
	Data []int
}

func main() {
	taro := MyData{
		"Taro",
		[]int{10, 20, 30},
	}
	fmt.Println(taro)
	rev(&taro)
	fmt.Println(taro)
}

func rev(md *MyData) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
