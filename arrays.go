package main

import "fmt"

func printArray(arr [3]int) {
	for i, v := range arr {
		fmt.Println(arr[i], v)
	}
	arr[0] = 100 //不改变主函数

}

func printArray2(arr *[3]int) {
	for i, v := range arr { //*arr 省略
		fmt.Println(arr[i], v)
	}
	arr[0] = 100 //改变主函数数组值

}

func main() {
	var arr1 [3]int
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//更好
	for i, v := range arr3 {
		fmt.Println(arr3[i], v)
	}

	printArray(arr1) //不能传入其他长度的数组 值传递拷贝数组
	fmt.Println(arr1)

	printArray2(&arr1) //不能传入其他长度的数组 值传递拷贝数组
	fmt.Println(arr1)  //改变了

}
