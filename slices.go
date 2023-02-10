package main

import "fmt"

// 不加长度就是slice
func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	//slice是对arr的视图
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//左开右闭
	s1 := arr[2:6]
	fmt.Println("arr[2:6]=", s1)
	s2 := arr[:6]
	fmt.Println("arr[:6]=", s2)
	s3 := arr[2:]
	fmt.Println("arr[2:]=", s3)
	s4 := arr[:]
	fmt.Println("arr[:]=", s4)

	updateSlice(s3)
	fmt.Println("arr[2:6]=", s1) // [100 3 4 5]
	fmt.Println(arr)             //[0 1 100 3 4 5 6 7 8 9]

	updateSlice(s2)
	fmt.Println("arr[:6]=", s2)
	fmt.Println(arr)
	fmt.Println("arr[:]=", s4)

	//tip:数组arr[6]变为slice arr[:]

	fmt.Println("reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arrtest := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arrtest)
	s11 := arrtest[2:6]
	//s11=[2 3 4 5],len(s11)=4,cap(s11)=6 cap指的是slice最大容量
	fmt.Printf("s11=%v,len(s11)=%d,cap(s11)=%d\n", s11, len(s11), cap(s11))
	s22 := s11[3:5] //重点越界len也能取到值!!!!Slice可以向后拓展不能向前拓展只要不超过底层cap
	//s22=[5 6],len(s22)=2,cap(s22)=3
	fmt.Printf("s22=%v,len(s22)=%d,cap(s22)=%d\n", s22, len(s22), cap(s22))

}
