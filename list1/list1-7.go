package main

import (
	"fmt"

	"github.com/taaag51/golang-handson/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
