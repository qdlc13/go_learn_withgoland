package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 外面不能用:= 不是全局变量 是包内部变量
var aa = 3
var ss = "kkk"
var (
	dd = 4
	gg = 4
)

// 包内常量 包内函数都能使用
const (
	tt = "yy"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}
func variableInitialValue() {
	var a, b int = 3, 5
	var s string = "abc"
	fmt.Printf("%d %d %q\n", a, b, s)
}
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "hi golang"
	fmt.Println(a, b, c, s)
}
func variableShorter() {
	a, b, c, s := 3, 4, true, "hi golang"
	fmt.Println(a, b, c, s)
}

// e^(i*PI)+1 = 0
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) //5

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) //(0+1.2246467991473515e-16i)接近0

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1) //(0+1.2246467991473515e-16i)
	//保留小数点后三位
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) //(0.000+0.000i)

}
func triangle() {
	a, b := 3, 4
	fmt.Println(calcTriangle(a, b))
}
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 常量定义
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4 //不定义类型 使用时看成文本替换
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c, tt)

}

// 枚举类型
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
	)
	//iota自增
	const (
		a = iota
		_
		b1
		c
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python)
	fmt.Println(a, b1, c)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("hello hi")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, dd, gg)
	//test
	//欧拉公式
	euler()
	triangle()
	consts()
	enums()

}
