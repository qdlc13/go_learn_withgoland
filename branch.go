package main

import (
	"fmt"
	"os"
)

func grade(score int) string {
	g := ""
	switch {
	case score >= 0 && score < 60:
		g = "F"
	case score >= 60 && score <= 100:
		g = "A"
	default:
		//panic函数需要参数所以使用Sprintf返回字符数组
		panic(fmt.Sprintf("wrong score %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"

	/* 等价 但是contents err作用域不同
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	*/
	//contents err只能在if中使用
	if contents, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//fmt.Println(grade(65),grade(-1))这样写在一起前面正确的也输出不了 一起恐慌)
	fmt.Println(grade(65))
	fmt.Println(grade(-1))
}
