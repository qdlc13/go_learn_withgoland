package main

import (
	"fmt"
	"go_learn_withgoland/test"
)

func main() {
	a := 2
	test.KK(func() { a++ })
	fmt.Println(a) //4
	//tests := []struct{ a, b, c int }{
	//	{3, 4, 5},
	//	{5, 12, 13},
	//	{8, 15, 17},
	//	{30000, 40000, 50000},
	//}
	//for a, tt := range tests {
	//	fmt.Println(a, tt)
	//	//if actual := calcTriangle(tt.a,tt.b);actual!=
	//}
}
