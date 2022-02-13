package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/taaag51/golang-handson/hello"
)

func main() {
	x := hello.Input("input data")
	ar := strings.Split(x, " ")
	t := 0

	for _, v := range ar {
		n, err := strconv.Atoi(v)
		if err != nil {
			goto err
		}
		t += n

	}

	fmt.Println("total:", t)
	return
err:
	fmt.Println("ERROR!!")
}
