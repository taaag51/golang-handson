package main

import (
	"fmt"
	"strconv"

	"github.com/taaag51/golang-handson/hello"
)

func main() {
	x := hello.Input("type a number")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数です。")
		} else {
			fmt.Println("奇数です。")
		}
	} else {
		fmt.Println("整数ではありません。")

	}
}
