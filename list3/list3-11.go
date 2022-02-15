package main

import (
	"fmt"
)

// MyData is structure.
type MyData struct {
	Name string
	Data []int
}

// PrintData is println all data.
func (md MyData) PrintData() {
	fmt.Println("*** MyData ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** end ***")
}

func main() {
	taro := MyData{
		"Hanako", []int{98, 76, 54, 32, 10},
	}
	taro.PrintData()
}
