package main

import (
	"fmt"
	"strconv"

	"github.com/taaag51/golang-handson/hello"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	fmt.Print(x + "は、")
	if n%2 == 0 {
		fmt.Println("偶数です。")
	} else {
		fmt.Println("奇数です。")
	}
}
