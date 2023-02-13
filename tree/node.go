package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

// 没有构造函数但可以有工厂函数
func creatNode(value int) *treeNode {
	return &treeNode{value: value} //c++中是错误返回了一个局部变量的地址 go中可用

}

// 方法不是函数
func (node treeNode) print() { //方法前的参数（使用者）也是传值类型
	fmt.Print(node.value) //值传递把node的内容拷贝一份给print函数打印
}
func (node treeNode) setValue(value int) {
	node.value = value
}
func (node *treeNode) setValue2(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " +
			"node. Ignored.")
		return
	}
	node.value = value //等价(*node).value = value
}
func main() {
	var root treeNode
	fmt.Println(root) //{0 <nil> <nil>}

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) //go特性 等价于其他语言root.right->left 无论是地址还是结构本身一律用.访问成员
	root.left.right = creatNode(2)

	nodes := []treeNode{ //结构体切片
		{value: 4},
		{},
		{5, &root, nil},
	}
	fmt.Println(nodes) //[{4 <nil> <nil>} {0 <nil> <nil>} {5 0xc000092060 <nil>}]
	root.print()       //3
	fmt.Println()
	root.right.left.setValue(5) //root.right.left地址此时被解析
	root.right.left.print()     // 0   setValue没有成功因为是值传递
	fmt.Println()

	root.right.left.setValue2(5)
	root.right.left.print() // 5   setValue2成功因为是指针
	fmt.Println()

	//无论定义的方法是值接收者还是很指针接收者，都可以用值类型调用
	//值类型会拷贝一份给值接收者，值类型会把地址传给指针接收者
	root.print()
	fmt.Println()
	root.setValue2(3)

	//无论定义的方法是值接收者还是很指针接收者，都可以用指针类型调用
	//指针类型会把地址所在的值传给值接收者，指针类型会把地址传给指针接收者
	pRoot := &root
	pRoot.print()
	fmt.Println()
	pRoot.setValue2(100)
	pRoot.print()
	fmt.Println()

	var testRoot *treeNode
	testRoot.setValue2(66)

	testRoot = &root
	testRoot.setValue2(500)
	testRoot.print() //500

}
