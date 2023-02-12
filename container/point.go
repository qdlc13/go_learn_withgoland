package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}
func swap2(a, b int) (int, int) {
	return b, a
}
func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	a1, b1 := 3, 4
	fmt.Println(swap2(a1, b1)) //更好
}
