package main

import (
	"fmt"
	"strconv"

	"github.com/taaag51/golang-handson/hello"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print(x + "は、")
	} else {
		return
	}
	switch {
	case n%2 == 0:
		fmt.Println("偶数です。")
	case n%2 == 1:
		fmt.Println("奇数です。")
	}
}
