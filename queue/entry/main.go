package main

import (
	"fmt"
	"go_learn_withgoland/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(4)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}
