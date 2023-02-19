package main

import (
	"fmt"
	"go_learn_withgoland/test"
)

func main() {
	a := 2
	test.KK(func() { a++ })
	fmt.Println(a) //4
}
