package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 没有构造函数但可以有工厂函数
func creatNode(value int) *treeNode {
	return &treeNode{value: value} //c++中是错误返回了一个局部变量的地址 go中可用

}

func (node treeNode) print() { //方法前的参数（使用者）也是传值类型
	fmt.Print(node.value) //值传递把node的内容拷贝一份给print函数打印
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()

}
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " +
			"node. Ignored.")
		return
	}
	node.value = value //等价(*node).value = value
}
func main() {
	root := treeNode{3, nil, nil}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) //go特性 等价于其他语言root.right->left 无论是地址还是结构本身一律用.访问成员
	root.left.right = creatNode(2)
	root.right.left.setValue(4)

	root.traverse()
}
