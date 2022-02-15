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
	taro := new(MyData)
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 5, 5)
	fmt.Println(taro)

}
