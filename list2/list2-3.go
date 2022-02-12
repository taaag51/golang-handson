package main

import (
	"fmt"
	"strconv"

	"github.com/taaag51/golang-handson/hello"
)

func main() {
	x := hello.Input("type a prise")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	p := float64(n)
	fmt.Println(int(p * 1.1))
}
