package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len(s)=%d,cap(s)=%d\n", len(s), cap(s))

}

func main() {
	fmt.Println("creating Slice")

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	s1 := arr[2:6]
	//s1=[2 3 4 5],len(s1)=4,cap(s1)=6 cap指的是slice最大容量
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := s1[3:5]
	//s2=[5 6],len(s2)=2,cap(s2)=3
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5 =", s3, s4, s5)
	fmt.Println("arr=", arr)
	// s4 and s5 no longer view arr.
	/*s4 s5的新的数组的view  添加元素超过cap系统重新分配更大的底层数组
	s3,s4,s5 = [5 6 10] [5 6 10 11] [5 6 10 11 12]
	arr= [0 1 2 3 4 5 6 10]
	*/

	//直接创建Slice
	var s []int //zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s) //cap装满就拓展一倍
		s = append(s, i*2+1)
	}
	fmt.Println(s)

	//创建Slice方式
	s6 := []int{2, 4, 6, 7, 3}
	printSlice(s6)
	//len为16
	s7 := make([]int, 16) //默认值为0
	printSlice(s7)
	//len为10 cap为32
	s8 := make([]int, 10, 32)
	printSlice(s8)

	fmt.Println("Copying slice")

}
