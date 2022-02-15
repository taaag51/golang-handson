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
	taro := MyData{"Taro", []int{10, 20, 30}}
	hanako := MyData{
		Name: "Hanako",
		Data: []int{10, 20, 30},
	}
	fmt.Println(taro)
	fmt.Println(hanako)
}
