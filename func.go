package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil

	default:
		//panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s ", op)
	}
}
func div(a, b int) (q, r int) {
	return a / b, a % b
}
func apply(op func(int, int) int, a, b int) int {
	//获取反射对象指针
	p := reflect.ValueOf(op).Pointer()
	//获取调用的函数名称
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

// 重写pow
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 65, "%")) //0 unsupported operation: %

	fmt.Println(eval(3, 65, "-"))
	if res, err := eval(3, 65, "%"); err != nil { //unsupported operation: %

		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	q, r := div(35, 6)
	fmt.Println(q, r)
	//Calling function main.pow with args(3,4)
	fmt.Println(apply(pow, 3, 4))
	//匿名函数//Calling function main.main.func1 with args(5,2)包名.函数名.匿名函数1
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 5, 2))

	fmt.Println(sum(1, 2, 3, 4))
}
